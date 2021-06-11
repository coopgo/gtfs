package gtfs

import (
	"archive/zip"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type opener interface {
	Open() (io.ReadCloser, error)
}

type namer interface {
	Name() string
}

type openNamer interface {
	opener
	namer
}

type closer interface {
	Close() error
}

type zipFile struct {
	f *zip.File
}

func (z zipFile) Open() (io.ReadCloser, error) {
	return z.f.Open()
}
func (z zipFile) Name() string {
	return z.f.Name
}
func (z zipFile) String() string {
	return z.f.Name
}

type regularFile struct {
	path string
	name string
}

func (r regularFile) Open() (io.ReadCloser, error) {
	return os.Open(r.path)
}
func (r regularFile) Name() string {
	return r.name
}
func (r regularFile) String() string {
	return r.path
}

type emptyCloser struct{}

func (emptyCloser) Close() error { return nil }

// TODO: Rename
func extract(path string) (closer, []openNamer, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, nil, err
	}

	var dir closer
	var files []openNamer

	if fi.IsDir() {
		filesInfo, err := ioutil.ReadDir(path)
		if err != nil {
			return nil, nil, err
		}
		files = make([]openNamer, len(filesInfo))
		for i, f := range filesInfo {
			files[i] = regularFile{
				path: filepath.Join(path, f.Name()),
				name: f.Name(),
			}
		}
		files = selectGtfsFiles(files)
		dir = emptyCloser{} // We do not need to close a directory since we only
		// work with the path
	} else {
		if filepath.Ext(path) != ".zip" {
			return nil, nil, errors.New("gtfs: invalid file extension")
		}

		r, err := zip.OpenReader(path)
		if err != nil {
			return nil, nil, err
		}

		files = make([]openNamer, len(r.File))
		for i, f := range r.File {
			files[i] = zipFile{f}
		}

		files = selectGtfsFiles(files)
		dir = r
	}
	return dir, files, err
}

func selectGtfsFiles(files []openNamer) []openNamer {
	var result []openNamer
	for _, f := range fileDepedanceOrder {
		if file, ok := contains(f, files); ok {
			result = append(result, file)
		}
	}
	return result
}

func contains(elem string, array []openNamer) (openNamer, bool) {
	for _, f := range array {
		if elem == f.Name() {
			return f, true
		}
	}
	return nil, false
}

type filename string

const (
	agency         filename = "agency.txt"
	levels         filename = "levels.txt"
	stops          filename = "stops.txt"
	transfers      filename = "transfers.txt"
	pathways       filename = "pathways.txt"
	calendar       filename = "calendar.txt"
	calendarDate   filename = "calendar_dates.txt"
	shape          filename = "shapes.txt"
	routes         filename = "routes.txt"
	trips          filename = "trips.txt"
	stoptimes      filename = "stop_times.txt"
	frequencies    filename = "frequencies.txt"
	fareAttributes filename = "fare_attributes.txt"
	fareRules      filename = "fare_rules.txt"
	feedInfo       filename = "feed_info.txt"
	translations   filename = "translations.txt"
	attributions   filename = "attributions.txt"
)

var fileDepedanceOrder []string = []string{
	string(agency),
	string(levels),
	string(stops),
	string(transfers),
	string(pathways),
	string(calendar),
	string(calendarDate),
	string(shape),
	string(routes),
	string(trips),
	string(stoptimes),
	string(frequencies),
	string(fareAttributes),
	string(fareRules),
	string(feedInfo),
	string(translations),
	string(attributions),
}
