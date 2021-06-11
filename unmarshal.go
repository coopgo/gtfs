// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2021-06-11 11:24:59.81057803 +0200 CEST m=+0.000136690

package gtfs

import "errors"

// ************************************************************************** //
// AgencySerializable GENERATED CODE
//
// Generated init and unmarshal function for AgencySerializable struct
// These functions sould be used to do a mapping between a slice of string
// and this struct.
// ************************************************************************** //

// init initalizes the values which have a gtfs empty value different from the
// golang zero value or if it's type is a custom one (for example Date).
func (item *AgencySerializable) init() {

}

// unmarshal takes a header slice which contains the gtfs tag of the field and a
// slice with contains the values of these fields.
// unmarshal makes the conversion of the d values in the correct type and check
// the presence of required fields (not conditionnal required fields).
func (item *AgencySerializable) unmarshal(h []string, d []string) error {
	var err error

	if len(h) != len(d) {
		return errors.New("header diff data")
	}

	item.init()

	for i := 0; i < len(h); i++ {
		switch h[i] {
		case "agency_id":
			item.Id, err = getString(d[i], false)
		case "agency_name":
			item.Name, err = getString(d[i], true)
		case "agency_url":
			item.Url, err = getString(d[i], true)
		case "agency_timezone":
			item.Timezone, err = getString(d[i], true)
		case "agency_lang":
			item.Lang, err = getString(d[i], false)
		case "agency_phone":
			item.Phone, err = getString(d[i], false)
		case "agency_fare_url":
			item.FareUrl, err = getString(d[i], false)
		case "agency_email":
			item.Email, err = getString(d[i], false)

		}

		if err != nil {
			return err
		}
	}
	return err
}

// ************************************************************************** //
// StopSerializable GENERATED CODE
//
// Generated init and unmarshal function for StopSerializable struct
// These functions sould be used to do a mapping between a slice of string
// and this struct.
// ************************************************************************** //

// init initalizes the values which have a gtfs empty value different from the
// golang zero value or if it's type is a custom one (for example Date).
func (item *StopSerializable) init() {

}

// unmarshal takes a header slice which contains the gtfs tag of the field and a
// slice with contains the values of these fields.
// unmarshal makes the conversion of the d values in the correct type and check
// the presence of required fields (not conditionnal required fields).
func (item *StopSerializable) unmarshal(h []string, d []string) error {
	var err error

	if len(h) != len(d) {
		return errors.New("header diff data")
	}

	item.init()

	for i := 0; i < len(h); i++ {
		switch h[i] {
		case "stop_id":
			item.Id, err = getString(d[i], true)
		case "stop_code":
			item.Code, err = getString(d[i], false)
		case "stop_name":
			item.Name, err = getString(d[i], false)
		case "stop_desc":
			item.Desc, err = getString(d[i], false)
		case "stop_lat":
			item.Lat, err = getFloat(d[i], false, 0)
		case "stop_lon":
			item.Long, err = getFloat(d[i], false, 0)
		case "stop_url":
			item.Url, err = getString(d[i], false)
		case "location_type":
			item.LocationType, err = getInt(d[i], false, 0)
		case "stop_timezone":
			item.Timezone, err = getString(d[i], false)
		case "wheelchair_boarding":
			item.WheelchairBoarding, err = getInt(d[i], false, 0)
		case "platform_code":
			item.PlatformCode, err = getString(d[i], false)
		case "parent_station":
			item.ParentStation, err = getString(d[i], false)
		case "level_id":
			item.LevelId, err = getString(d[i], false)
		case "zone_id":
			item.ZoneId, err = getString(d[i], false)

		}

		if err != nil {
			return err
		}
	}
	return err
}

// ************************************************************************** //
// RouteSerializable GENERATED CODE
//
// Generated init and unmarshal function for RouteSerializable struct
// These functions sould be used to do a mapping between a slice of string
// and this struct.
// ************************************************************************** //

// init initalizes the values which have a gtfs empty value different from the
// golang zero value or if it's type is a custom one (for example Date).
func (item *RouteSerializable) init() {

	item.ContinuousPickup = 1
	item.ContinuousDropOff = 1
}

// unmarshal takes a header slice which contains the gtfs tag of the field and a
// slice with contains the values of these fields.
// unmarshal makes the conversion of the d values in the correct type and check
// the presence of required fields (not conditionnal required fields).
func (item *RouteSerializable) unmarshal(h []string, d []string) error {
	var err error

	if len(h) != len(d) {
		return errors.New("header diff data")
	}

	item.init()

	for i := 0; i < len(h); i++ {
		switch h[i] {
		case "route_id":
			item.Id, err = getString(d[i], true)
		case "agency_id":
			item.AgencyId, err = getString(d[i], false)
		case "route_short_name":
			item.ShortName, err = getString(d[i], false)
		case "route_long_name":
			item.LongName, err = getString(d[i], false)
		case "route_desc":
			item.Desc, err = getString(d[i], false)
		case "route_type":
			item.Type, err = getInt(d[i], true, 0)
		case "route_url":
			item.Url, err = getString(d[i], false)
		case "route_color":
			item.Color, err = getString(d[i], false)
		case "route_text_color":
			item.TextColor, err = getString(d[i], false)
		case "route_sort_order":
			item.SortOrder, err = getInt(d[i], false, 0)
		case "continuous_pickup":
			item.ContinuousPickup, err = getInt(d[i], false, 1)
		case "continuous_drop_off":
			item.ContinuousDropOff, err = getInt(d[i], false, 1)

		}

		if err != nil {
			return err
		}
	}
	return err
}

// ************************************************************************** //
// TripSerializable GENERATED CODE
//
// Generated init and unmarshal function for TripSerializable struct
// These functions sould be used to do a mapping between a slice of string
// and this struct.
// ************************************************************************** //

// init initalizes the values which have a gtfs empty value different from the
// golang zero value or if it's type is a custom one (for example Date).
func (item *TripSerializable) init() {

}

// unmarshal takes a header slice which contains the gtfs tag of the field and a
// slice with contains the values of these fields.
// unmarshal makes the conversion of the d values in the correct type and check
// the presence of required fields (not conditionnal required fields).
func (item *TripSerializable) unmarshal(h []string, d []string) error {
	var err error

	if len(h) != len(d) {
		return errors.New("header diff data")
	}

	item.init()

	for i := 0; i < len(h); i++ {
		switch h[i] {
		case "trip_id":
			item.Id, err = getString(d[i], true)
		case "route_id":
			item.RouteId, err = getString(d[i], true)
		case "service_id":
			item.ServiceId, err = getString(d[i], true)
		case "block_id":
			item.BlockId, err = getString(d[i], false)
		case "shape_id":
			item.ShapeId, err = getString(d[i], false)
		case "trip_headsign":
			item.Headsign, err = getString(d[i], false)
		case "trip_short_name":
			item.ShortName, err = getString(d[i], false)
		case "direction_id":
			item.DirectionId, err = getInt(d[i], false, 0)
		case "wheelchair_accessible":
			item.WheelchairAccessible, err = getInt(d[i], false, 0)
		case "bikes_allowed":
			item.BikesAllowed, err = getInt(d[i], false, 0)

		}

		if err != nil {
			return err
		}
	}
	return err
}

// ************************************************************************** //
// StopTimeSerializable GENERATED CODE
//
// Generated init and unmarshal function for StopTimeSerializable struct
// These functions sould be used to do a mapping between a slice of string
// and this struct.
// ************************************************************************** //

// init initalizes the values which have a gtfs empty value different from the
// golang zero value or if it's type is a custom one (for example Date).
func (item *StopTimeSerializable) init() {

	item.ArrivalTime = TimeEmpty
	item.DepartureTime = TimeEmpty

	item.ContinuousPickup = 1
	item.ContinuousDropOff = 1

	item.Timepoint = 1
}

// unmarshal takes a header slice which contains the gtfs tag of the field and a
// slice with contains the values of these fields.
// unmarshal makes the conversion of the d values in the correct type and check
// the presence of required fields (not conditionnal required fields).
func (item *StopTimeSerializable) unmarshal(h []string, d []string) error {
	var err error

	if len(h) != len(d) {
		return errors.New("header diff data")
	}

	item.init()

	for i := 0; i < len(h); i++ {
		switch h[i] {
		case "trip_id":
			item.TripId, err = getString(d[i], true)
		case "stop_id":
			item.StopId, err = getString(d[i], true)
		case "arrival_time":
			item.ArrivalTime, err = getTime(d[i], false)
		case "departure_time":
			item.DepartureTime, err = getTime(d[i], false)
		case "stop_sequence":
			item.StopSequence, err = getInt(d[i], true, 0)
		case "stop_headsign":
			item.StopHeadsign, err = getString(d[i], false)
		case "pickup_type":
			item.PickupType, err = getInt(d[i], false, 0)
		case "drop_off_type":
			item.DropOffType, err = getInt(d[i], false, 0)
		case "continuous_pickup":
			item.ContinuousPickup, err = getInt(d[i], false, 1)
		case "continuous_drop_off":
			item.ContinuousDropOff, err = getInt(d[i], false, 1)
		case "":
			item.ShapeDistTraveled, err = getFloat(d[i], false, 0)
		case "timepoint":
			item.Timepoint, err = getInt(d[i], false, 1)

		}

		if err != nil {
			return err
		}
	}
	return err
}

// ************************************************************************** //
// CalendarDateSerializable GENERATED CODE
//
// Generated init and unmarshal function for CalendarDateSerializable struct
// These functions sould be used to do a mapping between a slice of string
// and this struct.
// ************************************************************************** //

// init initalizes the values which have a gtfs empty value different from the
// golang zero value or if it's type is a custom one (for example Date).
func (item *CalendarDateSerializable) init() {

	item.Date = DateEmpty
	item.Type = 1
}

// unmarshal takes a header slice which contains the gtfs tag of the field and a
// slice with contains the values of these fields.
// unmarshal makes the conversion of the d values in the correct type and check
// the presence of required fields (not conditionnal required fields).
func (item *CalendarDateSerializable) unmarshal(h []string, d []string) error {
	var err error

	if len(h) != len(d) {
		return errors.New("header diff data")
	}

	item.init()

	for i := 0; i < len(h); i++ {
		switch h[i] {
		case "service_id":
			item.ServiceId, err = getString(d[i], true)
		case "date":
			item.Date, err = getDate(d[i], true)
		case "exception_type":
			item.Type, err = getInt(d[i], true, 1)

		}

		if err != nil {
			return err
		}
	}
	return err
}

// ************************************************************************** //
// FareAttributeSerializable GENERATED CODE
//
// Generated init and unmarshal function for FareAttributeSerializable struct
// These functions sould be used to do a mapping between a slice of string
// and this struct.
// ************************************************************************** //

// init initalizes the values which have a gtfs empty value different from the
// golang zero value or if it's type is a custom one (for example Date).
func (item *FareAttributeSerializable) init() {

	item.Transfers = -1

}

// unmarshal takes a header slice which contains the gtfs tag of the field and a
// slice with contains the values of these fields.
// unmarshal makes the conversion of the d values in the correct type and check
// the presence of required fields (not conditionnal required fields).
func (item *FareAttributeSerializable) unmarshal(h []string, d []string) error {
	var err error

	if len(h) != len(d) {
		return errors.New("header diff data")
	}

	item.init()

	for i := 0; i < len(h); i++ {
		switch h[i] {
		case "fare_id":
			item.Id, err = getString(d[i], true)
		case "price":
			item.Price, err = getFloat(d[i], true, 0)
		case "currency_type":
			item.CurrencyType, err = getString(d[i], true)
		case "payment_method":
			item.PaymentMethod, err = getInt(d[i], true, 0)
		case "transfers":
			item.Transfers, err = getInt(d[i], true, -1)
		case "agency_id":
			item.AgencyId, err = getString(d[i], false)
		case "transfer_duration":
			item.TransferDuration, err = getInt(d[i], false, 0)

		}

		if err != nil {
			return err
		}
	}
	return err
}

// ************************************************************************** //
// FareRuleSerializable GENERATED CODE
//
// Generated init and unmarshal function for FareRuleSerializable struct
// These functions sould be used to do a mapping between a slice of string
// and this struct.
// ************************************************************************** //

// init initalizes the values which have a gtfs empty value different from the
// golang zero value or if it's type is a custom one (for example Date).
func (item *FareRuleSerializable) init() {

}

// unmarshal takes a header slice which contains the gtfs tag of the field and a
// slice with contains the values of these fields.
// unmarshal makes the conversion of the d values in the correct type and check
// the presence of required fields (not conditionnal required fields).
func (item *FareRuleSerializable) unmarshal(h []string, d []string) error {
	var err error

	if len(h) != len(d) {
		return errors.New("header diff data")
	}

	item.init()

	for i := 0; i < len(h); i++ {
		switch h[i] {
		case "fare_id":
			item.FareId, err = getString(d[i], true)
		case "route_id":
			item.RouteId, err = getString(d[i], false)
		case "origin_id":
			item.OriginId, err = getString(d[i], false)
		case "destination_id":
			item.DestinationId, err = getString(d[i], false)
		case "contains_id":
			item.ContainsId, err = getString(d[i], false)

		}

		if err != nil {
			return err
		}
	}
	return err
}

// ************************************************************************** //
// ShapeSerializable GENERATED CODE
//
// Generated init and unmarshal function for ShapeSerializable struct
// These functions sould be used to do a mapping between a slice of string
// and this struct.
// ************************************************************************** //

// init initalizes the values which have a gtfs empty value different from the
// golang zero value or if it's type is a custom one (for example Date).
func (item *ShapeSerializable) init() {

}

// unmarshal takes a header slice which contains the gtfs tag of the field and a
// slice with contains the values of these fields.
// unmarshal makes the conversion of the d values in the correct type and check
// the presence of required fields (not conditionnal required fields).
func (item *ShapeSerializable) unmarshal(h []string, d []string) error {
	var err error

	if len(h) != len(d) {
		return errors.New("header diff data")
	}

	item.init()

	for i := 0; i < len(h); i++ {
		switch h[i] {
		case "shape_id":
			item.ShapeId, err = getString(d[i], true)
		case "shape_pt_lat":
			item.Lat, err = getFloat(d[i], true, 0)
		case "shape_pt_lon":
			item.Long, err = getFloat(d[i], true, 0)
		case "shape_pt_sequence":
			item.Sequence, err = getInt(d[i], true, 0)
		case "shape_dist_traveled":
			item.DistTraveled, err = getFloat(d[i], false, 0)

		}

		if err != nil {
			return err
		}
	}
	return err
}

// ************************************************************************** //
// FrequencySerializable GENERATED CODE
//
// Generated init and unmarshal function for FrequencySerializable struct
// These functions sould be used to do a mapping between a slice of string
// and this struct.
// ************************************************************************** //

// init initalizes the values which have a gtfs empty value different from the
// golang zero value or if it's type is a custom one (for example Date).
func (item *FrequencySerializable) init() {

	item.StartTime = TimeEmpty
	item.EndTime = TimeEmpty

}

// unmarshal takes a header slice which contains the gtfs tag of the field and a
// slice with contains the values of these fields.
// unmarshal makes the conversion of the d values in the correct type and check
// the presence of required fields (not conditionnal required fields).
func (item *FrequencySerializable) unmarshal(h []string, d []string) error {
	var err error

	if len(h) != len(d) {
		return errors.New("header diff data")
	}

	item.init()

	for i := 0; i < len(h); i++ {
		switch h[i] {
		case "trip_id":
			item.TripId, err = getString(d[i], true)
		case "start_time":
			item.StartTime, err = getTime(d[i], true)
		case "end_time":
			item.EndTime, err = getTime(d[i], true)
		case "headway_secs":
			item.HeadwaySecs, err = getInt(d[i], true, 0)
		case "exact_times":
			item.ExactTimes, err = getInt(d[i], false, 0)

		}

		if err != nil {
			return err
		}
	}
	return err
}

// ************************************************************************** //
// TransferSerializable GENERATED CODE
//
// Generated init and unmarshal function for TransferSerializable struct
// These functions sould be used to do a mapping between a slice of string
// and this struct.
// ************************************************************************** //

// init initalizes the values which have a gtfs empty value different from the
// golang zero value or if it's type is a custom one (for example Date).
func (item *TransferSerializable) init() {

}

// unmarshal takes a header slice which contains the gtfs tag of the field and a
// slice with contains the values of these fields.
// unmarshal makes the conversion of the d values in the correct type and check
// the presence of required fields (not conditionnal required fields).
func (item *TransferSerializable) unmarshal(h []string, d []string) error {
	var err error

	if len(h) != len(d) {
		return errors.New("header diff data")
	}

	item.init()

	for i := 0; i < len(h); i++ {
		switch h[i] {
		case "from_stop_id":
			item.From, err = getString(d[i], true)
		case "to_stop_id":
			item.To, err = getString(d[i], true)
		case "transfer_type":
			item.Type, err = getInt(d[i], true, 0)
		case "min_transfer_time":
			item.MinTransferTime, err = getInt(d[i], false, 0)

		}

		if err != nil {
			return err
		}
	}
	return err
}

// ************************************************************************** //
// PathwaySerializable GENERATED CODE
//
// Generated init and unmarshal function for PathwaySerializable struct
// These functions sould be used to do a mapping between a slice of string
// and this struct.
// ************************************************************************** //

// init initalizes the values which have a gtfs empty value different from the
// golang zero value or if it's type is a custom one (for example Date).
func (item *PathwaySerializable) init() {

	item.Mode = 1

}

// unmarshal takes a header slice which contains the gtfs tag of the field and a
// slice with contains the values of these fields.
// unmarshal makes the conversion of the d values in the correct type and check
// the presence of required fields (not conditionnal required fields).
func (item *PathwaySerializable) unmarshal(h []string, d []string) error {
	var err error

	if len(h) != len(d) {
		return errors.New("header diff data")
	}

	item.init()

	for i := 0; i < len(h); i++ {
		switch h[i] {
		case "pathway_id":
			item.Id, err = getString(d[i], true)
		case "from_stop_id":
			item.From, err = getString(d[i], true)
		case "to_stop_id":
			item.To, err = getString(d[i], true)
		case "pathway_mode":
			item.Mode, err = getInt(d[i], true, 1)
		case "is_bidirectional":
			item.IsBidirectional, err = getInt(d[i], true, 0)
		case "length":
			item.Length, err = getFloat(d[i], false, 0)
		case "traversal_time":
			item.TraversalTime, err = getInt(d[i], false, 0)
		case "stair_count":
			item.StairCount, err = getInt(d[i], false, 0)
		case "max_slope":
			item.MaxSlope, err = getFloat(d[i], false, 0)
		case "min_width":
			item.MinWidth, err = getFloat(d[i], false, 0)
		case "signposted_as":
			item.SignpostedAs, err = getString(d[i], false)
		case "reversed_signposted_as":
			item.ReversedSignpostedAs, err = getString(d[i], false)

		}

		if err != nil {
			return err
		}
	}
	return err
}

// ************************************************************************** //
// LevelSerializable GENERATED CODE
//
// Generated init and unmarshal function for LevelSerializable struct
// These functions sould be used to do a mapping between a slice of string
// and this struct.
// ************************************************************************** //

// init initalizes the values which have a gtfs empty value different from the
// golang zero value or if it's type is a custom one (for example Date).
func (item *LevelSerializable) init() {

}

// unmarshal takes a header slice which contains the gtfs tag of the field and a
// slice with contains the values of these fields.
// unmarshal makes the conversion of the d values in the correct type and check
// the presence of required fields (not conditionnal required fields).
func (item *LevelSerializable) unmarshal(h []string, d []string) error {
	var err error

	if len(h) != len(d) {
		return errors.New("header diff data")
	}

	item.init()

	for i := 0; i < len(h); i++ {
		switch h[i] {
		case "level_id":
			item.Id, err = getString(d[i], true)
		case "level_index":
			item.Index, err = getFloat(d[i], true, 0)
		case "level_name":
			item.Name, err = getString(d[i], false)

		}

		if err != nil {
			return err
		}
	}
	return err
}

// ************************************************************************** //
// FeedInfoSerializable GENERATED CODE
//
// Generated init and unmarshal function for FeedInfoSerializable struct
// These functions sould be used to do a mapping between a slice of string
// and this struct.
// ************************************************************************** //

// init initalizes the values which have a gtfs empty value different from the
// golang zero value or if it's type is a custom one (for example Date).
func (item *FeedInfoSerializable) init() {

	item.StartDate = DateEmpty
	item.EndDate = DateEmpty

}

// unmarshal takes a header slice which contains the gtfs tag of the field and a
// slice with contains the values of these fields.
// unmarshal makes the conversion of the d values in the correct type and check
// the presence of required fields (not conditionnal required fields).
func (item *FeedInfoSerializable) unmarshal(h []string, d []string) error {
	var err error

	if len(h) != len(d) {
		return errors.New("header diff data")
	}

	item.init()

	for i := 0; i < len(h); i++ {
		switch h[i] {
		case "feed_publisher_name":
			item.PublisherName, err = getString(d[i], true)
		case "feed_publisher_url":
			item.PublisherUrl, err = getString(d[i], true)
		case "feed_lang":
			item.FeedLang, err = getString(d[i], true)
		case "default_lang":
			item.DefaultLang, err = getString(d[i], false)
		case "feed_start_date":
			item.StartDate, err = getDate(d[i], false)
		case "feed_end_date":
			item.EndDate, err = getDate(d[i], false)
		case "feed_version":
			item.Version, err = getString(d[i], false)
		case "feed_contact_email":
			item.ContactEmail, err = getString(d[i], false)
		case "feed_contact_url":
			item.ContactUrl, err = getString(d[i], false)

		}

		if err != nil {
			return err
		}
	}
	return err
}

// ************************************************************************** //
// TranslationSerializable GENERATED CODE
//
// Generated init and unmarshal function for TranslationSerializable struct
// These functions sould be used to do a mapping between a slice of string
// and this struct.
// ************************************************************************** //

// init initalizes the values which have a gtfs empty value different from the
// golang zero value or if it's type is a custom one (for example Date).
func (item *TranslationSerializable) init() {

}

// unmarshal takes a header slice which contains the gtfs tag of the field and a
// slice with contains the values of these fields.
// unmarshal makes the conversion of the d values in the correct type and check
// the presence of required fields (not conditionnal required fields).
func (item *TranslationSerializable) unmarshal(h []string, d []string) error {
	var err error

	if len(h) != len(d) {
		return errors.New("header diff data")
	}

	item.init()

	for i := 0; i < len(h); i++ {
		switch h[i] {
		case "table_name":
			item.TableName, err = getString(d[i], true)
		case "field_name":
			item.FieldName, err = getString(d[i], true)
		case "language":
			item.Language, err = getString(d[i], true)
		case "translation":
			item.Translation, err = getString(d[i], true)
		case "record_id":
			item.RecordId, err = getString(d[i], false)
		case "record_sub_id":
			item.RecordSubId, err = getString(d[i], false)
		case "field_value":
			item.FieldValue, err = getString(d[i], false)

		}

		if err != nil {
			return err
		}
	}
	return err
}

// ************************************************************************** //
// AttributionSerializable GENERATED CODE
//
// Generated init and unmarshal function for AttributionSerializable struct
// These functions sould be used to do a mapping between a slice of string
// and this struct.
// ************************************************************************** //

// init initalizes the values which have a gtfs empty value different from the
// golang zero value or if it's type is a custom one (for example Date).
func (item *AttributionSerializable) init() {

	item.IsProducer = 0
	item.IsOperator = 0
	item.IsAuthority = 0

}

// unmarshal takes a header slice which contains the gtfs tag of the field and a
// slice with contains the values of these fields.
// unmarshal makes the conversion of the d values in the correct type and check
// the presence of required fields (not conditionnal required fields).
func (item *AttributionSerializable) unmarshal(h []string, d []string) error {
	var err error

	if len(h) != len(d) {
		return errors.New("header diff data")
	}

	item.init()

	for i := 0; i < len(h); i++ {
		switch h[i] {
		case "":
			item.Id, err = getString(d[i], false)
		case "agency_id":
			item.AgencyId, err = getString(d[i], false)
		case "route_id":
			item.RouteId, err = getString(d[i], false)
		case "trip_id":
			item.TripId, err = getString(d[i], false)
		case "organization_name":
			item.OrganizationName, err = getString(d[i], true)
		case "is_producer":
			item.IsProducer, err = getInt(d[i], false, 0)
		case "is_operator":
			item.IsOperator, err = getInt(d[i], false, 0)
		case "is_authority":
			item.IsAuthority, err = getInt(d[i], false, 0)
		case "attribution_url":
			item.AttributionUrl, err = getString(d[i], false)
		case "attribution_email":
			item.AttributionEmail, err = getString(d[i], false)
		case "attribution_phone":
			item.AttributionPhone, err = getString(d[i], false)

		}

		if err != nil {
			return err
		}
	}
	return err
}