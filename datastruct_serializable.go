package gtfs

import (
	"errors"
	"fmt"
)

// This file contains the datastructure used by the gtfs library.
// There is two version for each structure in this library.
//
// A serializable version which is a 1-to-1 mapping to the gtfs file format as
// specified here: https://developers.google.com/transit/gtfs/reference
// (27/05/2021)
//
// A usable version which is an adapted version of the gtfs file format with
// linked reference between differents structs.
//
// This file contains the serializable version of the usable format of the
// datastructure.
//
// It should be used when as it names implied, you need serialize or persist the
// data.
//
// Note that the usable format can be unlinked to create a serializable format
// and that the serializable format can be linked to create a usable format.

// FeedSerializable contains different slices and maps that holds a GTFS dataset
// in the Serializable format.
type FeedSerializable struct {
	Agencies       map[string]*AgencySerializable
	Levels         map[string]*LevelSerializable
	Stops          map[string]*StopSerializable
	Transfers      []*TransferSerializable
	Pathways       map[string]*PathwaySerializable
	Calendar       map[string]*CalendarSerializable
	CalendarDates  map[string][]*CalendarDateSerializable
	Shapes         map[string][]*ShapeSerializable
	Routes         map[string]*RouteSerializable
	Trips          map[string]*TripSerializable
	StopTimes      map[string][]*StopTimeSerializable
	Frequencies    map[string]*FrequencySerializable
	FareAttributes map[string]*FareAttributeSerializable
	FareRules      []*FareRuleSerializable
	FeedInfo       *FeedInfoSerializable // There might be multiple FeedInfo per feed
	Translations   []*TranslationSerializable
	Attributions   []*AttributionSerializable
}

func (feed *FeedSerializable) init() {
	feed.Agencies = make(map[string]*AgencySerializable)
	feed.Levels = make(map[string]*LevelSerializable)
	feed.Stops = make(map[string]*StopSerializable)
	feed.Transfers = make([]*TransferSerializable, 0)
	feed.Pathways = make(map[string]*PathwaySerializable)
	feed.Calendar = make(map[string]*CalendarSerializable)
	feed.CalendarDates = make(map[string][]*CalendarDateSerializable)
	feed.Shapes = make(map[string][]*ShapeSerializable)
	feed.Routes = make(map[string]*RouteSerializable)
	feed.Trips = make(map[string]*TripSerializable)
	feed.StopTimes = make(map[string][]*StopTimeSerializable)
	feed.Frequencies = make(map[string]*FrequencySerializable)
	feed.FareAttributes = make(map[string]*FareAttributeSerializable)
	feed.FareRules = make([]*FareRuleSerializable, 0)
	feed.Translations = make([]*TranslationSerializable, 0)
	feed.Attributions = make([]*AttributionSerializable, 0)
}

// Validate apply a verification on the FeedSerializable to check if there is
// any error compared to the specification. If there is no error, Validate
// returns nil, otherwise it returns the first error it encounters.
func (feed *FeedSerializable) Validate() error {
	v := validator{}
	err := v.FeedSerializable(feed)
	if err != nil {
		return fmt.Errorf("feedserializable validation error: %w", err)
	}
	return nil
}

func (feed *FeedSerializable) String() string {

	stoptimes := 0
	for _, tr := range feed.StopTimes {
		stoptimes += len(tr)
	}

	calendarDates := 0
	for _, ser := range feed.CalendarDates {
		calendarDates += len(ser)
	}

	points := 0
	for _, shape := range feed.Shapes {
		points += len(shape)
	}

	return fmt.Sprintf("FeedSerializable{Agencies: %d, Stops: %d, Routes: %d, Trips: %d, Stoptimes: %d, Calendar: %d, CalendarDates: %d, FareAttributes: %d, FareRules: %d, Shapes: %d (%d), Frequencies: %d, Transfers: %d, Pathways: %d, Levels: %d, Translations: %d, Attributions: %d}",
		len(feed.Agencies), len(feed.Stops), len(feed.Routes), len(feed.Trips), stoptimes, len(feed.Calendar), calendarDates, len(feed.FareAttributes), len(feed.FareRules), len(feed.Shapes), points, len(feed.Frequencies), len(feed.Transfers), len(feed.Pathways), len(feed.Levels), len(feed.Translations), len(feed.Attributions))
}

func (feed *FeedSerializable) Link() (*Feed, error) {
	l := linker{}
	return l.Link(feed)
}

//go:generate go run gen.go

// All gtfs tags must be in the format :
//
// - gtfs:"tag" that implies that the field is optional and that it's empty
// value is the zero value (doesn't need to be set in the init method)
//
// - gtfs:"tag,required" implies that the tag is required, if the second arg
// is not "required", the tag is considered optional for the unmarshalling
//
// - gtfs:"tag,required,empty" implies that the tag is required like the last
// example. The "empty" tag is a special value that will be set in the generated
// init method.
//
// example:
// Route.ContinuousPickup `gtfs:"continuous_pickup,opt,1"`
//
// unmarshal.go ...
// func (item *RouteSerializable) init() {
//     ...
//     item.ContinuousPickup = 1
// }

type unmarshaller interface {
	unmarshal(h []string, d []string) error
}

type AgencySerializable struct {
	// Identifies a unique transit agency.
	// +conditionnally required - This field is required when the dataset
	// contains data for multiple transit agencies, otherwise it is optional.
	Id string `gtfs:"agency_id"`

	// Full name of the transit agency.
	// +required
	Name string `gtfs:"agency_name,required"`

	// URL of the transit agency.
	// +required
	Url string `gtfs:"agency_url,required"`

	// Timezone where the transit agency is located. If multiple agencies are
	// specified in the dataset, each must have the same agency_timezone.
	// +required
	Timezone string `gtfs:"agency_timezone,required"`

	// Primary language used by this transit agency.
	// +optional
	Lang string `gtfs:"agency_lang"`

	// A voice telephone number for the specified agency.
	// +optional
	Phone string `gtfs:"agency_phone"`

	// URL of a web page that allows a rider to purchase tickets or other fare
	// instruments for that agency online.
	// +optional
	FareUrl string `gtfs:"agency_fare_url"`

	// Email address actively monitored by the agency’s customer service
	// department.
	// +optional
	Email string `gtfs:"agency_email"`
}

type StopSerializable struct {

	// Identifies a stop, station, or station entrance.
	// +required
	Id string `gtfs:"stop_id,required"`

	// Short text or a number that identifies the location for riders.
	// +optional
	Code string `gtfs:"stop_code"`

	// Name of the location. Use a name that people will understand in the local
	// and tourist vernacular.
	// +conditionally required: https://developers.google.com/transit/gtfs/reference?hl=en#stopstxt
	Name string `gtfs:"stop_name"`

	// Description of the location that provides useful, quality information.
	// +optional
	Desc string `gtfs:"stop_desc"`

	// Latitude of the location.
	// +conditionally required: https://developers.google.com/transit/gtfs/reference?hl=en#stopstxt
	Lat float64 `gtfs:"stop_lat"`

	// Longitude of the location.
	// +conditionally required: https://developers.google.com/transit/gtfs/reference?hl=en#stopstxt
	Long float64 `gtfs:"stop_lon"`

	// URL of a web page about the location.
	// +optional
	Url string `gtfs:"stop_url"`

	// Type of the location.
	// See all values: https://developers.google.com/transit/gtfs/reference?hl=en#stopstxt
	// +optional
	LocationType int `gtfs:"location_type"`

	// Timezone of the location. If the location has a parent station, it
	// inherits the parent station’s timezone instead of applying its own.
	// +optional
	Timezone string `gtfs:"stop_timezone"`

	// Indicates whether wheelchair boardings are possible from the location.
	// See valid options: https://developers.google.com/transit/gtfs/reference?hl=en#stopstxt
	// +optional
	WheelchairBoarding int `gtfs:"wheelchair_boarding"`

	// Platform identifier for a platform stop (a stop belonging to a station).
	// +optional
	PlatformCode string `gtfs:"platform_code"`

	// Defines hierarchy between the different locations defined in stops.txt.
	// +conditionally required: https://developers.google.com/transit/gtfs/reference?hl=en#stopstxt
	ParentStation string `gtfs:"parent_station"`

	// Level of the location. The same level can be used by multiple unlinked
	// stations.
	// +optional
	LevelId string `gtfs:"level_id"`

	// Identifies the fare zone for a stop.
	// +conditionally required - This field is required if providing fare
	// information, otherwise it is optional. If this record represents a
	// station or station entrance, the ZoneId is ignored.
	ZoneId string `gtfs:"zone_id"`
}

type RouteSerializable struct {
	// Identifies a route.
	// +required
	Id string `gtfs:"route_id,required"`

	// Agency for the specified route.
	// +conditionally required - This field is required when the dataset
	// provides data for routes from more than one agency, otherwise it is
	// optional.
	AgencyId string `gtfs:"agency_id"`

	// Short name of a route.
	// +conditionally required - Either route_short_name or route_long_name must
	// be specified, or potentially both if appropriate.
	ShortName string `gtfs:"route_short_name"`

	// Full name of a route.
	// +conditionally required - Either route_short_name or route_long_name must
	// be specified, or potentially both if appropriate.
	LongName string `gtfs:"route_long_name"`

	// Description of a route that provides useful, quality information.
	// +optional
	Desc string `gtfs:"route_desc"`

	// Indicates the type of transportation used on a route.
	// See valid options: https://developers.google.com/transit/gtfs/reference?hl=en#routestxt
	// +required
	Type int `gtfs:"route_type,required"`

	// URL of a web page about the particular route. Should be different from
	// the Agency.Url value.
	// +optional
	Url string `gtfs:"route_url"`

	// Route color designation that matches public facing material.
	// +optional
	Color string `gtfs:"route_color"`

	// Legible color to use for text drawn against a background of route_color.
	// +optional
	TextColor string `gtfs:"route_text_color"`

	// Orders the routes in a way which is ideal for presentation to customers.
	// +optional
	SortOrder int `gtfs:"route_sort_order"`

	// Indicates whether a rider can board the transit vehicle anywhere along
	// the vehicle’s travel path.
	// See valid options: https://developers.google.com/transit/gtfs/reference?hl=en#routestxt
	// +optional
	ContinuousPickup int `gtfs:"continuous_pickup,opt,1"`

	// Indicates whether a rider can alight from the transit vehicle at any
	// point along the vehicle’s travel path.
	// See valid options: https://developers.google.com/transit/gtfs/reference?hl=en#routestxt
	// +optional
	ContinuousDropOff int `gtfs:"continuous_drop_off,opt,1"`
}

type TripSerializable struct {
	// Identifies a trip.
	// +required
	Id string `gtfs:"trip_id,required"`

	// Identifies a route.
	// +required
	RouteId string `gtfs:"route_id,required"`

	// Identifies a set of dates when service is available for one or more
	// routes.
	// +required
	ServiceId string `gtfs:"service_id,required"`

	// Identifies the block to which the trip belongs.
	// +optional
	BlockId string `gtfs:"block_id"`

	// Identifies a geospatial shape that describes the vehicle travel path for
	// a trip.
	// +conditionally required: This field is required if the trip has
	// continuous behavior defined, either at the route level or at the stop
	// time level. Otherwise, it's optional.
	ShapeId string `gtfs:"shape_id"`

	// Text that appears on signage identifying the trip's destination to
	// riders.
	// +optional
	Headsign string `gtfs:"trip_headsign"`

	// Public facing text used to identify the trip to riders, for instance, to
	// identify train numbers for commuter rail trips.
	// +optional
	ShortName string `gtfs:"trip_short_name"`

	// Indicates the direction of travel for a trip. This field is not used in
	// routing; it provides a way to separate trips by direction when publishing
	// time tables.
	// See valid options: https://developers.google.com/transit/gtfs/reference?hl=en#tripstxt
	// +optional
	DirectionId int `gtfs:"direction_id"`

	// Indicates wheelchair accessibility.
	// See valid options: https://developers.google.com/transit/gtfs/reference?hl=en#tripstxt
	// +optional
	WheelchairAccessible int `gtfs:"wheelchair_accessible"`

	// Indicates whether bikes are allowed.
	// See valid options: https://developers.google.com/transit/gtfs/reference?hl=en#tripstxt
	// +optional
	BikesAllowed int `gtfs:"bikes_allowed"`
}

type StopTimeSerializable struct {
	// Identifies a trip.
	// +required
	TripId string `gtfs:"trip_id,required"`

	// Identifies the serviced stop.
	// +required
	StopId string `gtfs:"stop_id,required"`

	// Arrival time at a specific stop for a specific trip on a route.
	// +conditionally required: https://developers.google.com/transit/gtfs/reference?hl=en#stop_timestxt
	ArrivalTime Time `gtfs:"arrival_time"`

	// Departure time from a specific stop for a specific trip on a route.
	// +conditionally required: https://developers.google.com/transit/gtfs/reference?hl=en#stop_timestxt
	DepartureTime Time `gtfs:"departure_time"`

	// Order of stops for a particular trip. The values must increase along the
	// trip but do not need to be consecutive.
	// +required
	StopSequence int `gtfs:"stop_sequence,required"`

	// Text that appears on signage identifying the trip's destination to
	// riders.
	// +optional
	StopHeadsign string `gtfs:"stop_headsign"`

	// Indicates pickup method.
	// See valid options: https://developers.google.com/transit/gtfs/reference?hl=en#stop_timestxt
	// +optional
	PickupType int `gtfs:"pickup_type"`

	// Indicates drop off method.
	// See valid options: https://developers.google.com/transit/gtfs/reference?hl=en#stop_timestxt
	// +optional
	DropOffType int `gtfs:"drop_off_type"`

	// Indicates whether a rider can board the transit vehicle at any point
	// along the vehicle’s travel path.
	// See valid options: https://developers.google.com/transit/gtfs/reference?hl=en#stop_timestxt
	// +optional
	ContinuousPickup int `gtfs:"continuous_pickup,opt,1"`

	// Indicates whether a rider can alight from the transit vehicle at any
	// point along the vehicle’s travel path.
	// See valid options: https://developers.google.com/transit/gtfs/reference?hl=en#stop_timestxt
	// +optional
	ContinuousDropOff int `gtfs:"continuous_drop_off,opt,1"`

	// Actual distance traveled along the associated shape, from the first stop
	// to the stop specified in this record. This field specifies how much of
	// the shape to draw between any two stops during a trip. Must be in the
	// same units used in Shape. Values used for ShapeDistTraveled must increase
	// along with StopSequence; they cannot be used to show reverse travel along
	// a route.
	// +optional
	ShapeDistTraveled float64 `csv:"shape_dist_traveled"`

	// Indicates if arrival and departure times for a stop are strictly adhered
	// to by the vehicle or if they are instead approximate and/or interpolated
	// times.
	// See valid options: https://developers.google.com/transit/gtfs/reference?hl=en#stop_timestxt
	// +optional
	Timepoint int `gtfs:"timepoint,opt,1"`
}

type CalendarSerializable struct {
	// Uniquely identifies a set of dates when service is available for one or
	// more routes. Each ServiceId value can appear at most once in a feed.
	// +required
	ServiceId string `gtfs:"service_id,required"`

	// Indicates whether the service operates on all specific days in the date
	// range specified by the StartDate and EndDate fields. This array starts
	// with Sunday and ends with Saturday.
	// +required
	Days [7]bool

	// Start service day for the service interval.
	// +required
	StartDate Date `gtfs:"start_date,required"`

	// End service day for the service interval. This service day is included in
	// the interval.
	// +required
	EndDate Date `gtfs:"end_date,required"`
}

// Order of the days as specified in the CalendarMeta.Days array.
var days [7]string = [7]string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

func (item *CalendarSerializable) init() {
	item.StartDate = DateEmpty
	item.EndDate = DateEmpty
}

func (item *CalendarSerializable) unmarshal(h []string, d []string) error {
	var err error
	var day string

	if len(h) != len(d) {
		return errors.New("header diff data")
	}

	item.init()

	for i := 0; i < len(h); i++ {
		switch h[i] {
		case "service_id":
			item.ServiceId, err = getString(d[i], true)
		case "start_date":
			item.StartDate, err = getDate(d[i], true)
		case "end_date":
			item.EndDate, err = getDate(d[i], true)
		case "sunday":
			day, err = getString(d[i], true)
			item.Days[0] = day == "1"
		case "monday":
			day, err = getString(d[i], true)
			item.Days[1] = day == "1"
		case "tuesday":
			day, err = getString(d[i], true)
			item.Days[2] = day == "1"
		case "wednesday":
			day, err = getString(d[i], true)
			item.Days[3] = day == "1"
		case "thursday":
			day, err = getString(d[i], true)
			item.Days[4] = day == "1"
		case "friday":
			day, err = getString(d[i], true)
			item.Days[5] = day == "1"
		case "saturday":
			day, err = getString(d[i], true)
			item.Days[6] = day == "1"
		}
		if err != nil {
			return err
		}
	}

	return err
}

type CalendarDateSerializable struct {
	// Identifies a set of dates when a service exception occurs for one or more
	// routes.
	// +required
	ServiceId string `gtfs:"service_id,required"`

	// Date when service exception occurs.
	// +required
	Date Date `gtfs:"date,required"`

	// Indicates whether service is available on the date specified in the date field.
	// See valid options: https://developers.google.com/transit/gtfs/reference?hl=en#calendar_datestxt
	// +required
	Type int `gtfs:"exception_type,required,1"`
}

type FareAttributeSerializable struct {
	// Identifies a fare class.
	// +required
	Id string `gtfs:"fare_id,required"`

	// Fare price, in the unit specified by currency_type.
	// +required
	Price float64 `gtfs:"price,required"`

	// Currency used to pay the fare.
	// +required
	CurrencyType string `gtfs:"currency_type,required"`

	// Indicates when the fare must be paid.
	// See valid options: https://developers.google.com/transit/gtfs/reference?hl=en#fare_attributestxt
	// +required
	PaymentMethod int `gtfs:"payment_method,required"`

	// Indicates the number of transfers permitted on this fare.
	// See valid options:
	// 0 - No transfers permitted on this fare.
	// 1 - Riders may transfer once.
	// 2 - Riders may transfer twice.
	// -1 - Unlimited transfers are permitted. This differs from the original
	// specification: https://developers.google.com/transit/gtfs/reference?hl=en#fare_attributestxt
	//
	// +required
	Transfers int `gtfs:"transfers,required,-1"`

	// Identifies the relevant agency for a fare.
	// +conditionnaly required - This field is required for datasets with
	// multiple agencies, otherwise it is optional.
	AgencyId string `gtfs:"agency_id"`

	// Length of time in seconds before a transfer expires.
	// +optional
	TransferDuration int `gtfs:"transfer_duration"`
}

type FareRuleSerializable struct {

	// Identifies a fare class.
	// +required
	FareId string `gtfs:"fare_id,required"`

	// Identifies a route associated with the fare class.
	// +optional
	RouteId string `gtfs:"route_id"`

	// Identifies an origin zone.
	// +optional
	OriginId string `gtfs:"origin_id"`

	// Identifies a destination zone.
	// +optional
	DestinationId string `gtfs:"destination_id"`

	// Identifies the zones that a rider will enter while using a given fare
	// class.
	// +optional
	ContainsId string `gtfs:"contains_id"`
}

type ShapeSerializable struct {
	// Identifies a shape.
	// +required
	ShapeId string `gtfs:"shape_id,required"`

	// Latitude of a shape point.
	// +required
	Lat float64 `gtfs:"shape_pt_lat,required"`

	// Longitude of a shape point.
	// +required
	Long float64 `gtfs:"shape_pt_lon,required"`

	// Sequence in which the shape points connect to form the shape. Values must
	// increase along the trip but do not need to be consecutive.
	// +required
	Sequence int `gtfs:"shape_pt_sequence,required"`

	// Actual distance traveled along the shape from the first shape point to
	// the point specified in this record.
	// +optional
	DistTraveled float64 `gtfs:"shape_dist_traveled"`
}

type FrequencySerializable struct {
	// Identifies a trip to which the specified headway of service applies.
	// +required
	TripId string `gtfs:"trip_id,required"`

	// Time at which the first vehicle departs from the first stop of the trip
	// with the specified headway.
	// +required
	StartTime Time `gtfs:"start_time,required"`

	// Time at which service changes to a different headway (or ceases) at the
	// first stop in the trip.
	// +required
	EndTime Time `gtfs:"end_time,required"`

	// Time, in seconds, between departures from the same stop (headway) for the
	// trip, during the time interval specified by start_time and end_time.
	// +required
	HeadwaySecs int `gtfs:"headway_secs,required"`

	// Indicates the type of service for a trip. See the file description for
	// more information.
	// See valid options: https://developers.google.com/transit/gtfs/reference?hl=en#frequenciestxt
	// +optional
	ExactTimes int `gtfs:"exact_times"`
}

type TransferSerializable struct {
	// Identifies a stop or station where a connection between routes begins.
	// If this field refers to a station, the transfer rule applies to all its
	// child stops.
	// +required
	From string `gtfs:"from_stop_id,required"`

	// Identifies a stop or station where a connection between routes ends.
	// If this field refers to a station, the transfer rule applies to all child
	// stops.
	// +required
	To string `gtfs:"to_stop_id,required"`

	// Indicates the type of connection for the specified
	// (from_stop_id, to_stop_id) pair.
	// See valid options: https://developers.google.com/transit/gtfs/reference?hl=en#transferstxt
	// +required
	Type int `gtfs:"transfer_type,required"`

	// Amount of time, in seconds, that must be available to permit a transfer
	// between routes at the specified stops.
	// +optional
	MinTransferTime int `gtfs:"min_transfer_time"`
}

type PathwaySerializable struct {

	// The pathway_id field contains an ID that uniquely identifies the pathway.
	// The pathway_id is used by systems as an internal identifier of this
	// record (e.g., primary key in database), and therefore the pathway_id
	// must be dataset unique.
	// +required
	Id string `gtfs:"pathway_id,required"`

	// Location at which the pathway begins.
	// +required
	From string `gtfs:"from_stop_id,required"`

	// Location at which the pathway ends.
	// +required
	To string `gtfs:"to_stop_id,required"`

	// Type of pathway between the specified (from_stop_id, to_stop_id) pair.
	// See valid values: https://developers.google.com/transit/gtfs/reference?hl=en#pathwaystxt
	// +required
	Mode int `gtfs:"pathway_mode,required,1"`

	// Indicates in which direction the pathway can be used.
	// +required
	// FIXME: change this into a bool
	IsBidirectional int `gtfs:"is_bidirectional,required"`

	// Horizontal length in meters of the pathway from the origin location to
	// the destination location.
	// +optional
	Length float64 `gtfs:"length"`

	// Average time in seconds needed to walk through the pathway from the
	// origin location to the destination location.
	// +optional
	TraversalTime int `gtfs:"traversal_time"`

	// Number of stairs of the pathway.
	// +optional
	StairCount int `gtfs:"stair_count"`

	// Maximum slope ratio of the pathway.
	// See valid values: https://developers.google.com/transit/gtfs/reference?hl=en#pathwaystxt
	// +optional
	MaxSlope float64 `gtfs:"max_slope"`

	// Minimum width of the pathway in meters.
	// +optional
	MinWidth float64 `gtfs:"min_width"`

	// String of text from physical signage visible to transit riders.
	// +optional
	SignpostedAs string `gtfs:"signposted_as"`

	// Same than the signposted_as field, but when the pathways is used backward
	// ie from the destination to the origin.
	// +optional
	ReversedSignpostedAs string `gtfs:"reversed_signposted_as"`
}

type LevelSerializable struct {
	// Id of the level that can be referenced from stops.txt.
	// +required
	Id string `gtfs:"level_id,required"`

	// Numeric index of the level that indicates relative position of this level
	// in relation to other levels (levels with higher indices are assumed to be
	// located above levels with lower indices).
	// +required
	Index float64 `gtfs:"level_index,required"`

	// Optional name of the level (that matches level lettering/numbering used
	//inside the building or the station).
	// +optional
	Name string `gtfs:"level_name"`
}

type FeedInfoSerializable struct {
	// Full name of the organization that publishes the dataset.
	// +required
	PublisherName string `gtfs:"feed_publisher_name,required"`

	// URL of the dataset publishing organization's website.
	// +required
	PublisherUrl string `gtfs:"feed_publisher_url,required"`

	// Default language for the text in this dataset.
	// See more: https://developers.google.com/transit/gtfs/reference#feed_infotxt
	// +required
	FeedLang string `gtfs:"feed_lang,required"`

	// Defines the language used when the data consumer doesn’t know the
	// language of the rider. It's often defined as en, English.
	// +optional
	DefaultLang string `gtfs:"default_lang"`

	// The dataset provides complete and reliable schedule information for
	// service in the period from the beginning of the StartDate day to the end
	// of the EndDate day.
	// +optional
	StartDate Date `gtfs:"feed_start_date"`

	// The dataset provides complete and reliable schedule information for
	// service in the period from the beginning of the StartDate day to the end
	// of the EndDate day.
	// +optional
	EndDate Date `gtfs:"feed_end_date"`

	// String that indicates the current version of their GTFS dataset.
	// GTFS-consuming applications can display this value to help dataset
	// publishers determine whether the latest dataset has been incorporated.*
	// +optional
	Version string `gtfs:"feed_version"`

	// Email address for communication regarding the GTFS dataset and data
	// publishing practices. ContactEmail is a technical contact for
	// GTFS-consuming applications. Provide customer service contact information
	// through agency.txt.
	// +optional
	ContactEmail string `gtfs:"feed_contact_email"`

	// URL for contact information, a web-form, support desk, or other tools for
	// communication regarding the GTFS dataset and data publishing practices.
	// ContactUrl is a technical contact for GTFS-consuming applications.
	// Provide customer service contact information through agency.txt.
	// +optional
	ContactUrl string `gtfs:"feed_contact_url"`
}

type TranslationSerializable struct {
	// Defines the dataset table that contains the field to be translated.
	// See valid values: https://developers.google.com/transit/gtfs/reference#translationstxt
	TableName string `gtfs:"table_name,required"`

	// Provides the name of the field to be translated.
	// +required
	FieldName string `gtfs:"field_name,required"`

	// Provides the language of translation.
	// See more: https://developers.google.com/transit/gtfs/reference#translationstxt
	// +required
	Language string `gtfs:"language,required"`

	// Provides the translated value for the specified FieldName.
	// +required
	Translation string `gtfs:"translation,required"`

	// Defines the record that corresponds to the field to be translated.
	// The value in RecordId needs to be a main ID from a dataset table.
	// See more: https://developers.google.com/transit/gtfs/reference#translationstxt
	//
	// +conditionnaly required - The following conditions determine how this
	// field can be used:
	// - Forbidden if TableName equals FeedInfo.
	// - Forbidden if FieldValue is defined.
	// - Required if FieldValue is empty.
	RecordId string `gtfs:"record_id"`

	// Helps to translate the record that contains the field when the table
	// referenced in RecordId doesn’t have a unique ID.
	// This means that this field is only useful when the TableName is
	// stop_times. (RecordSubId should be the StopSequence then.)
	//
	// +conditionnaly required - The following conditions determine how this
	// field can be used:
	// - Forbidden if TableName equals FeedInfo.
	// - Forbidden if FieldValue is defined.
	// - Required if TableName equals stop_times and RecordId is defined.
	RecordSubId string `gtfs:"record_sub_id"`

	// Instead of using RecordId and RecordSubId to define which record needs
	// to be translated, FieldValue can be used to define the value for
	// translation. When used, the translation is applied when the field
	// identified by TableName and FieldName contains the exact same value
	// defined in FieldValue.
	//
	// The field must exactly match the value defined in FieldValue. If only a
	// subset of the value matches FieldValue, the translation isn't applied.
	//
	// If two translation rules match the same record, one with FieldValue and
	// the other one with RecordId, then the rule with RecordId is the one
	// that needs to be used.
	//
	// +conditionnaly required - The following conditions determine how this
	// field can be used:
	// - Forbidden if TableName equals FeedInfo.
	// - Forbidden if RecordId is defined.
	// - Required if RecordId is empty.
	FieldValue string `gtfs:"field_value"`
}

type AttributionSerializable struct {

	// Identifies an attribution for the dataset, or a subset of it. This field
	// is useful for translations.
	// +optional
	Id string

	// The agency to which the attribution applies. Multiple attributions can
	// apply to the same trip. AgencyId, RouteId and TripId are exclusive, there
	// can only be one field set per attribution.
	// +optional
	AgencyId string `gtfs:"agency_id"`

	// The route to which the attribution applies. Multiple attributions can
	// apply to the same trip. AgencyId, RouteId and TripId are exclusive, there
	// can only be one field set per attribution.
	// +optional
	RouteId string `gtfs:"route_id"`

	// The trip to which the attribution applies. Multiple attributions can
	// apply to the same trip. AgencyId, RouteId and TripId are exclusive, there
	// can only be one field set per attribution.
	// +optional
	TripId string `gtfs:"trip_id"`

	// The name of the organization that the dataset is attributed to.
	// +required
	OrganizationName string `gtfs:"organization_name,required"`

	// The role of the organization is producer.
	//
	// Allowed values include the following:
	// - 0 or empty: Organization doesn’t have this role.
	// - 1: Organization does have this role.
	//
	// +conditionnaly required - At least one of the fields, either IsProducer,
	// IsOperator, or IsAuthority, must be set at 1.
	IsProducer int `gtfs:"is_producer,cond,0"`

	// The role of the organization is operator.
	//
	// Allowed values include the following:
	// - 0 or empty: Organization doesn’t have this role.
	// - 1: Organization does have this role.
	//
	// +conditionnaly required - At least one of the fields, either IsProducer,
	// IsOperator, or IsAuthority, must be set at 1.
	IsOperator int `gtfs:"is_operator,cond,0"`

	// The role of the organization is authority.
	//
	// Allowed values include the following:
	// - 0 or empty: Organization doesn’t have this role.
	// - 1: Organization does have this role.
	//
	// +conditionnaly required - At least one of the fields, either IsProducer,
	// IsOperator, or IsAuthority, must be set at 1.
	IsAuthority int `gtfs:"is_authority,cond,0"`

	// The URL of the organization.
	// +optional
	AttributionUrl string `gtfs:"attribution_url"`

	// The email of the organization.
	// +optional
	AttributionEmail string `gtfs:"attribution_email"`

	// The phone number of the organization.
	// +optional
	AttributionPhone string `gtfs:"attribution_phone"`
}
