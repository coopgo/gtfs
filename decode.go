package gtfs

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"sync"
)

// Parser is the struct that contains all the option to parse GTFS files.
type Parser struct {
	Validation bool
}

func (p Parser) Print() {
	fmt.Printf("Parser%+v\n", p)
}

func NoValidation(p *Parser) {
	p.Validation = false
}

// ParserOption is a function that sets a certain config on a Parser.
//
// This is part of the self referential functions design.
// See more: https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
type ParserOption func(*Parser)

// NewParser creates a new custom parser.
//
// Use self referential functions design to configure.
// See more: https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
func NewParser(opts ...ParserOption) Parser {
	parser := Parser{
		Validation: true,
	}

	for _, opt := range opts {
		opt(&parser)
	}

	return parser
}

func (parser *Parser) Load(path string) (*Feed, error) {

	// We generate our OpenNamers
	dir, files, err := extract(path)
	if err != nil {
		return nil, fmt.Errorf("gtfs: %w", err)
	}
	defer dir.Close()

	feedSerializable, err := parser.unmarshal(files)
	if err != nil {
		return nil, fmt.Errorf("gtfs: %w", err)
	}

	feed, err := feedSerializable.Link()
	if err != nil {
		return nil, fmt.Errorf("gtfs: %w", err)
	}

	if parser.Validation {
		if err := feed.Validate(); err != nil {
			return nil, fmt.Errorf("gtfs: %w", err)
		}
	}

	return feed, nil
}

func (parser *Parser) Unmarshal(path string) (*FeedSerializable, error) {

	// We generate our OpenNamers
	dir, files, err := extract(path)
	if err != nil {
		return nil, fmt.Errorf("gtfs: %w", err)
	}
	defer dir.Close()

	feed, err := parser.unmarshal(files)
	if err != nil {
		return nil, fmt.Errorf("gtfs: %w", err)
	}

	if parser.Validation {
		if err := feed.Validate(); err != nil {
			return nil, fmt.Errorf("gtfs: %w", err)
		}
	}

	return feed, nil
}

func (parser *Parser) unmarshal(files []openNamer) (*FeedSerializable, error) {
	feed := &FeedSerializable{}
	feed.init()

	// Synchronization init
	var wg sync.WaitGroup
	wg.Add(len(files))
	done := make(chan struct{})
	errs := make(chan error)

	for _, file := range files {
		go func(file openNamer) {
			d := decoder{
				feed: feed,
				wg:   &wg,
				done: done,
				errs: errs,
			}
			d.exec(file)
		}(file)
	}

	// When all the files have finished to be parsed, we can close the errs
	// channel.
	go func() {
		wg.Wait()
		close(errs)
	}()

	// We check the errors, if we found one error, we close the done channel to
	// stop every subroutine.
	for err := range errs {
		if err != nil {
			close(done)
			return nil, err
		}
	}

	return feed, nil
}

// decoder handles a reader and a worker to read an openNamer and exports it
// into a feed specific field.
type decoder struct {
	feed *FeedSerializable
	wg   *sync.WaitGroup
	done chan struct{}
	errs chan error
}

func (d *decoder) exec(f openNamer) {
	defer d.wg.Done()

	r := reader{
		done: d.done,
		errs: d.errs,
	}
	err := r.init(f)
	if err != nil {
		d.errs <- err
		return
	}
	data := r.read()

	w := worker{
		file:   f,
		header: r.header,
		done:   d.done,
		errs:   d.errs,
	}
	items := w.exec(data)

	for item := range items {
		select {
		case <-d.done:
			return
		default:
			if err := d.add(f, item); err != nil {
				d.errs <- err
				return
			}
		}
	}
}

func (d *decoder) add(file namer, item unmarshaller) error {
	switch filename(file.Name()) {
	case agency:
		i, ok := item.(*AgencySerializable)
		if !ok {
			return errors.New("invalid conversion")
		}
		d.feed.Agencies[i.Id] = i
	case levels:
		i, ok := item.(*LevelSerializable)
		if !ok {
			return errors.New("invalid conversion")
		}
		d.feed.Levels[i.Id] = i
	case stops:
		i, ok := item.(*StopSerializable)
		if !ok {
			return errors.New("invalid conversion")
		}
		d.feed.Stops[i.Id] = i
	case transfers:
		i, ok := item.(*TransferSerializable)
		if !ok {
			return errors.New("invalid conversion")
		}
		d.feed.Transfers = append(d.feed.Transfers, i)
	case pathways:
		i, ok := item.(*PathwaySerializable)
		if !ok {
			return errors.New("invalid conversion")
		}
		d.feed.Pathways[i.Id] = i
	case calendar:
		i, ok := item.(*CalendarSerializable)
		if !ok {
			return errors.New("invalid conversion")
		}
		d.feed.Calendar[i.ServiceId] = i
	case calendarDate:
		i, ok := item.(*CalendarDateSerializable)
		if !ok {
			return errors.New("invalid conversion")
		}
		d.feed.CalendarDates[i.ServiceId] = append(d.feed.CalendarDates[i.ServiceId], i)
	case shape:
		i, ok := item.(*ShapeSerializable)
		if !ok {
			return errors.New("invalid conversion")
		}
		d.feed.Shapes[i.ShapeId] = append(d.feed.Shapes[i.ShapeId], i)
	case routes:
		i, ok := item.(*RouteSerializable)
		if !ok {
			return errors.New("invalid conversion")
		}
		d.feed.Routes[i.Id] = i
	case trips:
		i, ok := item.(*TripSerializable)
		if !ok {
			return errors.New("invalid conversion")
		}
		d.feed.Trips[i.Id] = i
	case stoptimes:
		i, ok := item.(*StopTimeSerializable)
		if !ok {
			return errors.New("invalid conversion")
		}
		d.feed.StopTimes[i.TripId] = append(d.feed.StopTimes[i.TripId], i)
	case frequencies:
		i, ok := item.(*FrequencySerializable)
		if !ok {
			return errors.New("invalid conversion")
		}
		d.feed.Frequencies[i.TripId] = i
	case fareAttributes:
		i, ok := item.(*FareAttributeSerializable)
		if !ok {
			return errors.New("invalid conversion")
		}
		d.feed.FareAttributes[i.Id] = i
	case fareRules:
		i, ok := item.(*FareRuleSerializable)
		if !ok {
			return errors.New("invalid conversion")
		}
		d.feed.FareRules = append(d.feed.FareRules, i)
	case feedInfo:
		i, ok := item.(*FeedInfoSerializable)
		if !ok {
			return errors.New("invalid conversion")
		}
		d.feed.FeedInfo = i
	case translations:
		i, ok := item.(*TranslationSerializable)
		if !ok {
			return errors.New("invalid conversion")
		}
		d.feed.Translations = append(d.feed.Translations, i)
	case attributions:
		i, ok := item.(*AttributionSerializable)
		if !ok {
			return errors.New("invalid conversion")
		}
		d.feed.Attributions = append(d.feed.Attributions, i)
	default:
		// This should not be possible.
		return errors.New("internal error, invalid filename while adding to feed")
	}
	return nil
}

// reader job is to takes an openNamer as input, open it and read the data while
// doing a simple parsing to export arrays of string.
type reader struct {
	readcloser io.ReadCloser
	csvReader  *csv.Reader
	header     []string
	name       string

	done chan struct{}
	errs chan error
}

func (r *reader) init(f openNamer) error {
	var err error

	r.readcloser, err = f.Open()
	if err != nil {
		return err
	}

	r.name = f.Name()

	r.csvReader = csv.NewReader(r.readcloser)
	r.csvReader.FieldsPerRecord = -1
	r.csvReader.TrimLeadingSpace = true

	record, err := r.csvReader.Read()
	if err != nil {
		if err == io.EOF {
			return fmt.Errorf("empty file error")
		} else {
			return err
		}
	}

	r.header = record
	return nil
}

func (r *reader) close() error {
	r.csvReader = nil
	return r.readcloser.Close()
}

func (r *reader) read() chan []string {
	out := make(chan []string, 1000)

	go func() {
		defer r.close()
		defer close(out)
		for {
			select {
			case <-r.done:
				return
			default:
				record, err := r.csvReader.Read()
				if err != nil {
					if err == io.EOF {
						return
					} else {
						r.errs <- fmt.Errorf("reader error in %s: %w", r.name, err)
						return
					}
				}

				// We check that the size of the record is the same as the
				// header, if it's inferior we augment it, if it's superior we
				// return an err
				if len(record) < len(r.header) {
					t := make([]string, len(r.header))
					copy(t, record)
					record = t
				}
				if len(record) > len(r.header) {
					err := fmt.Errorf("record bigger than header")
					r.errs <- fmt.Errorf("reader error in %s: %w", r.name, err)
					return
				}

				out <- record
			}
		}
	}()
	return out
}

// worker job is to create the gtfs structs from the data it receives as input
// (normally from the reader).
type worker struct {
	file   namer
	header []string

	done chan struct{}
	errs chan error
}

// exec creates an item corresponding to the correct file, unmarshal the data
// coming from the in channel and puts their results in different output
// channel.
func (w *worker) exec(in chan []string) chan unmarshaller {
	out := make(chan unmarshaller, 1000)

	go func() {
		defer close(out)

		for data := range in {
			item, err := generator{}.newUnmarshaller(filename(w.file.Name()))
			if err != nil {
				w.errs <- err
				return
			}

			if err := item.unmarshal(w.header, data); err != nil {
				w.errs <- err
				return
			}

			select {
			case out <- item:
			case <-w.done:
				return
			}
		}

	}()

	return out
}

type generator struct{}

// TODO: Add new structs here
func (g generator) newUnmarshaller(f filename) (unmarshaller, error) {
	switch f {
	case agency:
		return &AgencySerializable{}, nil
	case levels:
		return &LevelSerializable{}, nil
	case stops:
		return &StopSerializable{}, nil
	case transfers:
		return &TransferSerializable{}, nil
	case pathways:
		return &PathwaySerializable{}, nil
	case calendar:
		return &CalendarSerializable{}, nil
	case calendarDate:
		return &CalendarDateSerializable{}, nil
	case shape:
		return &ShapeSerializable{}, nil
	case routes:
		return &RouteSerializable{}, nil
	case trips:
		return &TripSerializable{}, nil
	case stoptimes:
		return &StopTimeSerializable{}, nil
	case frequencies:
		return &FrequencySerializable{}, nil
	case fareAttributes:
		return &FareAttributeSerializable{}, nil
	case fareRules:
		return &FareRuleSerializable{}, nil
	default:
		return nil, errors.New("invalid file")
	}
}
