package gtfs

/*
	FIXME: I don't like the implementation done here.
	I would probably need to refactor it but it's not a priority right now.

*/

type validator struct{}

func (v validator) Feed(feed *Feed) error {
	done := make(chan struct{})
	items := v.reader(feed, done)
	errs := v.worker(feed, done, items)

	for err := range errs {
		if err != nil { // failsafe just in case
			close(done)
			return err
		}
	}

	return nil
}

func (v validator) FeedSerializable(feed *FeedSerializable) error {

	done := make(chan struct{})
	items := v.readerSerializable(feed, done)
	errs := v.workerSerializable(feed, done, items)

	for err := range errs {
		if err != nil { // failsafe just in case
			close(done)
			return err
		}
	}

	return nil
}

func (v *validator) reader(feed *Feed, done chan struct{}) chan validater {
	items := make(chan validater, 100)

	go func() {
		defer close(items)

		// we read here
		for _, item := range feed.Agencies {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, item := range feed.Stops {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, item := range feed.Routes {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, item := range feed.Trips {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, item := range feed.StopTimes() {
			select {
			case <-done:
				return
			default:
				items <- item
			}

		}

		for _, item := range feed.Services {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, item := range feed.FareAttributes {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, item := range feed.FareRules() {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, item := range feed.Shapes {
			select {
			case <-done:
				return
			default:
				items <- item
			}

		}

		for _, item := range feed.Frequencies {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, item := range feed.Transfers {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, item := range feed.Pathways {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, item := range feed.Levels {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		/*
			if feed.FeedInfo != nil {
				items <- feed.FeedInfo
			}

			for _, item := range feed.Translations {
				select {
				case <-done:
					return
				default:
					items <- item
				}
			}

			for _, item := range feed.Attributions {
				select {
				case <-done:
					return
				default:
					items <- item
				}
			}
		*/

	}()

	return items
}

func (v *validator) worker(feed *Feed, done chan struct{}, items chan validater) chan error {
	errs := make(chan error, 10)

	go func() {
		defer close(errs)
		for item := range items {
			select {
			case <-done:
				return
			default:
				if err := item.validate(feed); err != nil {
					errs <- err
				}

			}
		}
	}()

	return errs
}

func (v *validator) readerSerializable(feed *FeedSerializable, done chan struct{}) chan serializableValidater {
	items := make(chan serializableValidater, 100)
	go func() {
		defer close(items)

		for _, item := range feed.Agencies {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, item := range feed.Stops {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, item := range feed.Routes {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, item := range feed.Trips {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, trip := range feed.StopTimes {
			for _, item := range trip {
				select {
				case <-done:
					return
				default:
					items <- item
				}
			}
		}

		for _, item := range feed.Calendar {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, service := range feed.CalendarDates {
			for _, item := range service {
				select {
				case <-done:
					return
				default:
					items <- item
				}
			}
		}

		for _, item := range feed.FareAttributes {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, item := range feed.FareRules {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, shape := range feed.Shapes {
			for _, item := range shape {
				select {
				case <-done:
					return
				default:
					items <- item
				}
			}
		}

		for _, item := range feed.Frequencies {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, item := range feed.Transfers {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, item := range feed.Pathways {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, item := range feed.Levels {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		if feed.FeedInfo != nil {
			items <- feed.FeedInfo
		}

		for _, item := range feed.Translations {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

		for _, item := range feed.Attributions {
			select {
			case <-done:
				return
			default:
				items <- item
			}
		}

	}()
	return items
}

func (v *validator) workerSerializable(feed *FeedSerializable, done chan struct{}, items chan serializableValidater) chan error {
	errs := make(chan error, 10)

	go func() {
		defer close(errs)
		for item := range items {
			select {
			case <-done:
				return
			default:
				if err := item.validate(feed); err != nil {
					errs <- err
				}

			}
		}
	}()

	return errs
}

/*
	Gtfs object validation


*/

type validater interface {
	validate(feed *Feed) error
}

type serializableValidater interface {
	validate(feed *FeedSerializable) error
}

func (item *Agency) validate(feed *Feed) error {
	if len(feed.Agencies) > 1 && item.Id == "" {
		return &ConditionnalRequirementError{
			Struct:    "AgencySerializable",
			Field:     "Id",
			Condition: "This field is required when the dataset contains data for multiple transit agencies, otherwise it is optional.",
		}
	}

	if item.Name == "" {
		return &RequirementError{
			Struct: "AgencySerializable",
			Field:  "Name",
		}
	}

	if item.Url == "" {
		return &RequirementError{
			Struct: "AgencySerializable",
			Field:  "Url",
		}
	}

	if item.Timezone == "" {
		return &RequirementError{
			Struct: "AgencySerializable",
			Field:  "Timezone",
		}
	}

	return nil
}

func (item *AgencySerializable) validate(feed *FeedSerializable) error {
	if len(feed.Agencies) > 1 && item.Id == "" {
		return &ConditionnalRequirementError{
			Struct:    "AgencySerializable",
			Field:     "Id",
			Condition: "This field is required when the dataset contains data for multiple transit agencies, otherwise it is optional.",
		}
	}

	if item.Name == "" {
		return &RequirementError{
			Struct: "AgencySerializable",
			Field:  "Name",
		}
	}

	if item.Url == "" {
		return &RequirementError{
			Struct: "AgencySerializable",
			Field:  "Url",
		}
	}

	if item.Timezone == "" {
		return &RequirementError{
			Struct: "AgencySerializable",
			Field:  "Timezone",
		}
	}

	return nil
}

func (item *Stop) validate(feed *Feed) error {
	if item.Id == "" {
		return &RequirementError{
			Struct: "StopSerializable",
			Field:  "Id",
		}
	}

	if in(item.LocationType, 0, 1, 2) {
		if item.Name == "" {
			return &ConditionnalRequirementError{
				Struct:    "StopSerializable",
				Field:     "Name",
				Condition: "Required for locations which are stops (location_type=0), stations (location_type=1) or entrances/exits (location_type=2).",
			}
		}

		// This should be ok since the point 0,0 in in the atlantic ocean.
		if item.Lat == 0 {
			return &ConditionnalRequirementError{
				Struct:    "StopSerializable",
				Field:     "Lat",
				Condition: "Required for locations which are stops (location_type=0), stations (location_type=1) or entrances/exits (location_type=2).",
			}
		}

		if item.Long == 0 {
			return &ConditionnalRequirementError{
				Struct:    "StopSerializable",
				Field:     "Lat",
				Condition: "Required for locations which are stops (location_type=0), stations (location_type=1) or entrances/exits (location_type=2).",
			}
		}
	}
	/* Does not work with the Sample Feed from google
	if len(feed.FareRules()) != 0 && item.Zone == nil {
		return &ConditionnalRequirementError{
			Struct:    "StopSerializable",
			Field:     "ZoneId",
			Condition: "This field is required if providing fare information, otherwise it is optional.",
		}
	}
	*/

	if in(item.LocationType, 2, 3, 4) && item.ParentStation == nil {
		return &ConditionnalRequirementError{
			Struct:    "StopSerializable",
			Field:     "ParentStation",
			Condition: "Required for locations which are entrances (location_type=2), generic nodes (location_type=3) or boarding areas (location_type=4).",
		}
	} else if item.LocationType == 1 && item.ParentStation != nil {
		return &ConditionnalRequirementError{
			Struct:    "StopSerializable",
			Field:     "ParentStation",
			Condition: "Forbidden for stations (location_type=1).",
		}
	}

	if !in(item.LocationType, 0, 1, 2, 3, 4) {
		return &InvalidValueError{
			Struct: "StopSerializable",
			Field:  "LocationType",
			Value:  item.LocationType,
		}
	}

	if !in(item.WheelchairBoarding, 0, 1, 2) {
		return &InvalidValueError{
			Struct: "StopSerializable",
			Field:  "WheelchairBoarding",
			Value:  item.WheelchairBoarding,
		}
	}

	if item.Level != nil {
		if _, ok := feed.Levels[item.Level.Id]; ok {
			return &ReferenceError{
				Struct:      "StopSerializable",
				Field:       "LevelId",
				Destination: "LevelSerializable",
				Value:       item.Level.Id,
			}
		}
	}

	return nil
}

func (item *StopSerializable) validate(feed *FeedSerializable) error {

	if item.Id == "" {
		return &RequirementError{
			Struct: "StopSerializable",
			Field:  "Id",
		}
	}

	if in(item.LocationType, 0, 1, 2) {
		if item.Name == "" {
			return &ConditionnalRequirementError{
				Struct:    "StopSerializable",
				Field:     "Name",
				Condition: "Required for locations which are stops (location_type=0), stations (location_type=1) or entrances/exits (location_type=2).",
			}
		}

		// This should be ok since the point 0,0 in in the atlantic ocean.
		if item.Lat == 0 {
			return &ConditionnalRequirementError{
				Struct:    "StopSerializable",
				Field:     "Lat",
				Condition: "Required for locations which are stops (location_type=0), stations (location_type=1) or entrances/exits (location_type=2).",
			}
		}

		if item.Long == 0 {
			return &ConditionnalRequirementError{
				Struct:    "StopSerializable",
				Field:     "Lat",
				Condition: "Required for locations which are stops (location_type=0), stations (location_type=1) or entrances/exits (location_type=2).",
			}
		}
	}

	/* Does not work with Sample Feed from google
	if len(feed.FareRules) != 0 && item.ZoneId == "" {
		return &ConditionnalRequirementError{
			Struct:    "StopSerializable",
			Field:     "ZoneId",
			Condition: "This field is required if providing fare information, otherwise it is optional.",
		}
	}
	*/

	if in(item.LocationType, 2, 3, 4) && item.ParentStation == "" {
		return &ConditionnalRequirementError{
			Struct:    "StopSerializable",
			Field:     "ParentStation",
			Condition: "Required for locations which are entrances (location_type=2), generic nodes (location_type=3) or boarding areas (location_type=4).",
		}
	} else if item.LocationType == 1 && item.ParentStation != "" {
		return &ConditionnalRequirementError{
			Struct:    "StopSerializable",
			Field:     "ParentStation",
			Condition: "Forbidden for stations (location_type=1).",
		}
	}

	if !in(item.LocationType, 0, 1, 2, 3, 4) {
		return &InvalidValueError{
			Struct: "StopSerializable",
			Field:  "LocationType",
			Value:  item.LocationType,
		}
	}

	if item.ParentStation != "" {
		if _, ok := feed.Stops[item.ParentStation]; !ok {
			return &ReferenceError{
				Struct:      "StopSerializable",
				Field:       "ParentStation",
				Destination: "StopSerializable",
				Value:       item.ParentStation,
			}
		}
	}

	if !in(item.WheelchairBoarding, 0, 1, 2) {
		return &InvalidValueError{
			Struct: "StopSerializable",
			Field:  "WheelchairBoarding",
			Value:  item.WheelchairBoarding,
		}
	}

	if item.LevelId != "" {
		if _, ok := feed.Levels[item.LevelId]; ok {
			return &ReferenceError{
				Struct:      "StopSerializable",
				Field:       "LevelId",
				Destination: "LevelSerializable",
				Value:       item.LevelId,
			}
		}
	}

	return nil
}

func (item *Route) validate(feed *Feed) error {
	if item.Id == "" {
		return &RequirementError{
			Struct: "RouteSerializable",
			Field:  "Id",
		}
	}

	// Might be invalid
	if len(feed.Agencies) > 1 && item.Agency == nil {
		return &ConditionnalRequirementError{
			Struct:    "RouteSerializable",
			Field:     "AgencyId",
			Condition: "This field is required when the dataset provides data for routes from more than one agency, otherwise it is optional.",
		}
	}

	if item.Agency != nil {
		if _, ok := feed.Agencies[item.Agency.Id]; !ok {
			return &ReferenceError{
				Struct:      "RouteSerializable",
				Field:       "AgencyId",
				Destination: "AgencySerializable",
				Value:       item.Agency.Id,
			}
		}
	}

	if item.ShortName == "" && item.LongName == "" {
		return &ConditionnalRequirementError{
			Struct:    "RouteSerializable",
			Field:     "ShortName",
			Condition: "Either ShortName or LongName must be specified, or potentially both if appropriate.",
		}
	}

	if !in(item.Type, 0, 1, 2, 3, 4, 5, 6, 7, 11, 12) {
		return &InvalidValueError{
			Struct: "RouteSerializable",
			Field:  "Type",
			Value:  item.Type,
		}
	}

	if !in(item.ContinuousPickup, 0, 1, 2, 3) {
		return &InvalidValueError{
			Struct: "RouteSerializable",
			Field:  "ContinuousPickup",
			Value:  item.ContinuousPickup,
		}
	}

	if !in(item.ContinuousDropOff, 0, 1, 2, 3) {
		return &InvalidValueError{
			Struct: "RouteSerializable",
			Field:  "ContinuousDropOff",
			Value:  item.ContinuousDropOff,
		}
	}

	return nil
}

func (item *RouteSerializable) validate(feed *FeedSerializable) error {

	if item.Id == "" {
		return &RequirementError{
			Struct: "RouteSerializable",
			Field:  "Id",
		}
	}

	// Might be invalid
	if len(feed.Agencies) > 1 && item.AgencyId == "" {
		return &ConditionnalRequirementError{
			Struct:    "RouteSerializable",
			Field:     "AgencyId",
			Condition: "This field is required when the dataset provides data for routes from more than one agency, otherwise it is optional.",
		}
	}

	if item.AgencyId != "" {
		if _, ok := feed.Agencies[item.AgencyId]; !ok {
			return &ReferenceError{
				Struct:      "RouteSerializable",
				Field:       "AgencyId",
				Destination: "AgencySerializable",
				Value:       item.AgencyId,
			}
		}
	}

	if item.ShortName == "" && item.LongName == "" {
		return &ConditionnalRequirementError{
			Struct:    "RouteSerializable",
			Field:     "ShortName",
			Condition: "Either ShortName or LongName must be specified, or potentially both if appropriate.",
		}
	}

	if !in(item.Type, 0, 1, 2, 3, 4, 5, 6, 7, 11, 12) {
		return &InvalidValueError{
			Struct: "RouteSerializable",
			Field:  "Type",
			Value:  item.Type,
		}
	}

	if !in(item.ContinuousPickup, 0, 1, 2, 3) {
		return &InvalidValueError{
			Struct: "RouteSerializable",
			Field:  "ContinuousPickup",
			Value:  item.ContinuousPickup,
		}
	}

	if !in(item.ContinuousDropOff, 0, 1, 2, 3) {
		return &InvalidValueError{
			Struct: "RouteSerializable",
			Field:  "ContinuousDropOff",
			Value:  item.ContinuousDropOff,
		}
	}

	return nil
}

func (item *Trip) validate(feed *Feed) error {
	_, ok := feed.Routes[item.Route.Id]
	if item.Route == nil {
		return &RequirementError{
			Struct: "TripSerializable",
			Field:  "RouteId",
		}
	} else {
		if !ok {
			return &ReferenceError{
				Struct:      "TripSerializable",
				Field:       "RouteId",
				Destination: "RouteSerializable",
				Value:       item.Route.Id,
			}
		}
	}

	if item.Service == nil {
		return &RequirementError{
			Struct: "TripSerializable",
			Field:  "ServiceId",
		}
	} else {
		if _, ok := feed.Services[item.Service.Id]; !ok {
			return &ReferenceError{
				Struct:      "TripSerializable",
				Field:       "ServiceId",
				Destination: "Service",
				Value:       item.Service.Id,
			}
		}
	}

	if item.Id == "" {
		return &RequirementError{
			Struct: "TripSerializable",
			Field:  "Id",
		}
	}

	if !in(item.DirectionId, 0, 1) {
		return &InvalidValueError{
			Struct: "TripSerializable",
			Field:  "DirectionId",
			Value:  item.DirectionId,
		}
	}

	// Might be wrong
	/*
		routeContinousBehaviour := (route.ContinuousPickup == 0 || route.ContinuousDropOff == 0)
		stoptimeContinuousBehaviour := false

		stoptimes := feed.StopTimes[item.Id]
		for _, stoptime := range stoptimes {
			if stoptime.ContinuousPickup == 0 || stoptime.ContinuousDropOff == 0 {
				stoptimeContinuousBehaviour = true
				break
			}
		}

		if (routeContinousBehaviour || stoptimeContinuousBehaviour) && item.ShapeId == "" {
			return &ConditionnalRequirementError{
				Struct:    "TripSerializable",
				Field:     "ShapeId",
				Condition: "This field is required if the trip has continuous behavior defined, either at the route level or at the stop time level.",
			}
		}
	*/

	if item.Shape != nil {
		if _, ok := feed.Shapes[item.Shape.Id]; !ok {
			return &ReferenceError{
				Struct:      "TripSerializable",
				Field:       "ShapeId",
				Destination: "ShapeSerializable",
				Value:       item.Shape.Id,
			}
		}
	}

	if !in(item.WheelchairAccessible, 0, 1, 2) {
		return &InvalidValueError{
			Struct: "TripSerializable",
			Field:  "WheelchairAccessible",
			Value:  item.WheelchairAccessible,
		}
	}

	if !in(item.BikesAllowed, 0, 1, 2) {
		return &InvalidValueError{
			Struct: "TripSerializable",
			Field:  "BikesAllowed",
			Value:  item.BikesAllowed,
		}
	}

	return nil
}

func (item *TripSerializable) validate(feed *FeedSerializable) error {
	_, ok := feed.Routes[item.RouteId]
	if item.RouteId == "" {
		return &RequirementError{
			Struct: "TripSerializable",
			Field:  "RouteId",
		}
	} else {
		if !ok {
			return &ReferenceError{
				Struct:      "TripSerializable",
				Field:       "RouteId",
				Destination: "RouteSerializable",
				Value:       item.RouteId,
			}
		}
	}

	if item.ServiceId == "" {
		return &RequirementError{
			Struct: "TripSerializable",
			Field:  "ServiceId",
		}
	} else {
		_, ok1 := feed.Calendar[item.ServiceId]
		_, ok2 := feed.CalendarDates[item.ServiceId]
		if !ok1 && !ok2 {
			return &ReferenceError{
				Struct:      "TripSerializable",
				Field:       "ServiceId",
				Destination: "CalendarSerializable or CalendarDateSerializable",
				Value:       item.ServiceId,
			}
		}
	}

	if item.Id == "" {
		return &RequirementError{
			Struct: "TripSerializable",
			Field:  "Id",
		}
	}

	if !in(item.DirectionId, 0, 1) {
		return &InvalidValueError{
			Struct: "TripSerializable",
			Field:  "DirectionId",
			Value:  item.DirectionId,
		}
	}

	// Might be wrong
	/*
		routeContinousBehaviour := (route.ContinuousPickup == 0 || route.ContinuousDropOff == 0)
		stoptimeContinuousBehaviour := false

		stoptimes := feed.StopTimes[item.Id]
		for _, stoptime := range stoptimes {
			if stoptime.ContinuousPickup == 0 || stoptime.ContinuousDropOff == 0 {
				stoptimeContinuousBehaviour = true
				break
			}
		}

		if (routeContinousBehaviour || stoptimeContinuousBehaviour) && item.ShapeId == "" {
			return &ConditionnalRequirementError{
				Struct:    "TripSerializable",
				Field:     "ShapeId",
				Condition: "This field is required if the trip has continuous behavior defined, either at the route level or at the stop time level.",
			}
		}
	*/

	if item.ShapeId != "" {
		if _, ok := feed.Shapes[item.ShapeId]; !ok {
			return &ReferenceError{
				Struct:      "TripSerializable",
				Field:       "ShapeId",
				Destination: "ShapeSerializable",
				Value:       item.ShapeId,
			}
		}
	}

	if !in(item.WheelchairAccessible, 0, 1, 2) {
		return &InvalidValueError{
			Struct: "TripSerializable",
			Field:  "WheelchairAccessible",
			Value:  item.WheelchairAccessible,
		}
	}

	if !in(item.BikesAllowed, 0, 1, 2) {
		return &InvalidValueError{
			Struct: "TripSerializable",
			Field:  "BikesAllowed",
			Value:  item.BikesAllowed,
		}
	}

	return nil
}

func (item *StopTime) validate(feed *Feed) error {
	if item.Trip == nil {
		return &RequirementError{
			Struct: "StopTimeSerializable",
			Field:  "TripId",
		}
	} else {
		if _, ok := feed.Trips[item.Trip.Id]; !ok {
			return &ReferenceError{
				Struct:      "StopTimeSerializable",
				Field:       "TripId",
				Destination: "TripSerializable",
				Value:       item.Trip.Id,
			}
		}
	}

	if item.Timepoint == 1 && item.ArrivalTime.IsEmpty() {
		return &ConditionnalRequirementError{
			Struct:    "StopTimeSerializable",
			Field:     "ArrivalTime",
			Condition: "Provide arrival times for all stops that are time points. ",
		}
	}

	if item.Timepoint == 1 && item.DepartureTime.IsEmpty() {
		return &ConditionnalRequirementError{
			Struct:    "StopTimeSerializable",
			Field:     "DepartureTime",
			Condition: "Provide departure times for all stops that are time points. ",
		}
	}

	if item.Stop == nil {
		return &RequirementError{
			Struct: "StopTimeSerializable",
			Field:  "StopId",
		}
	} else {
		if _, ok := feed.Stops[item.Stop.Id]; !ok {
			return &ReferenceError{
				Struct:      "StopTimeSerializable",
				Field:       "StopId",
				Destination: "StopSerializable",
				Value:       item.Stop.Id,
			}
		}
	}

	if item.StopSequence < 0 {
		return &InvalidValueError{
			Struct: "StopTimeSerializable",
			Field:  "StopSequence",
			Value:  item.StopSequence,
		}
	}

	if !in(item.PickupType, 0, 1, 2, 3) {
		return &InvalidValueError{
			Struct: "StopTimeSerializable",
			Field:  "PickupType",
			Value:  item.PickupType,
		}
	}

	if !in(item.DropOffType, 0, 1, 2, 3) {
		return &InvalidValueError{
			Struct: "StopTimeSerializable",
			Field:  "DropOffType",
			Value:  item.DropOffType,
		}
	}

	if !in(item.ContinuousPickup, 0, 1, 2, 3) {
		return &InvalidValueError{
			Struct: "StopTimeSerializable",
			Field:  "ContinuousPickup",
			Value:  item.ContinuousPickup,
		}
	}

	if !in(item.ContinuousDropOff, 0, 1, 2, 3) {
		return &InvalidValueError{
			Struct: "StopTimeSerializable",
			Field:  "ContinuousDropOff",
			Value:  item.ContinuousDropOff,
		}
	}

	if item.ShapeDistTraveled < 0 {
		return &InvalidValueError{
			Struct: "StopTimeSerializable",
			Field:  "ShapeDistTraveled",
			Value:  item.ShapeDistTraveled,
		}
	}

	if !in(item.Timepoint, 0, 1) {
		return &InvalidValueError{
			Struct: "StopTimeSerializable",
			Field:  "Timepoint",
			Value:  item.Timepoint,
		}
	}

	return nil
}

func (item *StopTimeSerializable) validate(feed *FeedSerializable) error {

	if item.TripId == "" {
		return &RequirementError{
			Struct: "StopTimeSerializable",
			Field:  "TripId",
		}
	} else {
		if _, ok := feed.Trips[item.TripId]; !ok {
			return &ReferenceError{
				Struct:      "StopTimeSerializable",
				Field:       "TripId",
				Destination: "TripSerializable",
				Value:       item.TripId,
			}
		}
	}

	if item.Timepoint == 1 && item.ArrivalTime.IsEmpty() {
		return &ConditionnalRequirementError{
			Struct:    "StopTimeSerializable",
			Field:     "ArrivalTime",
			Condition: "Provide arrival times for all stops that are time points. ",
		}
	}

	if item.Timepoint == 1 && item.DepartureTime.IsEmpty() {
		return &ConditionnalRequirementError{
			Struct:    "StopTimeSerializable",
			Field:     "DepartureTime",
			Condition: "Provide departure times for all stops that are time points. ",
		}
	}

	if item.StopId == "" {
		return &RequirementError{
			Struct: "StopTimeSerializable",
			Field:  "StopId",
		}
	} else {
		if _, ok := feed.Stops[item.StopId]; !ok {
			return &ReferenceError{
				Struct:      "StopTimeSerializable",
				Field:       "StopId",
				Destination: "StopSerializable",
				Value:       item.StopId,
			}
		}
	}

	if item.StopSequence < 0 {
		return &InvalidValueError{
			Struct: "StopTimeSerializable",
			Field:  "StopSequence",
			Value:  item.StopSequence,
		}
	}

	if !in(item.PickupType, 0, 1, 2, 3) {
		return &InvalidValueError{
			Struct: "StopTimeSerializable",
			Field:  "PickupType",
			Value:  item.PickupType,
		}
	}

	if !in(item.DropOffType, 0, 1, 2, 3) {
		return &InvalidValueError{
			Struct: "StopTimeSerializable",
			Field:  "DropOffType",
			Value:  item.DropOffType,
		}
	}

	if !in(item.ContinuousPickup, 0, 1, 2, 3) {
		return &InvalidValueError{
			Struct: "StopTimeSerializable",
			Field:  "ContinuousPickup",
			Value:  item.ContinuousPickup,
		}
	}

	if !in(item.ContinuousDropOff, 0, 1, 2, 3) {
		return &InvalidValueError{
			Struct: "StopTimeSerializable",
			Field:  "ContinuousDropOff",
			Value:  item.ContinuousDropOff,
		}
	}

	if item.ShapeDistTraveled < 0 {
		return &InvalidValueError{
			Struct: "StopTimeSerializable",
			Field:  "ShapeDistTraveled",
			Value:  item.ShapeDistTraveled,
		}
	}

	if !in(item.Timepoint, 0, 1) {
		return &InvalidValueError{
			Struct: "StopTimeSerializable",
			Field:  "Timepoint",
			Value:  item.Timepoint,
		}
	}

	return nil
}

func (item *Service) validate(feed *Feed) error {

	if item.Id == "" {
		return &RequirementError{
			Struct: "Service",
			Field:  "Id",
		}
	}

	if item.Calendar == nil && len(item.Exceptions) == 0 {
		return &ConditionnalRequirementError{
			Struct:    "Service",
			Field:     "Calendar or Exceptions ",
			Condition: "A service must have a Calendar or at least one Exception.",
		}
	}

	if item.Calendar != nil {
		if err := item.Calendar.validate(feed); err != nil {
			return err
		}
	}

	for _, exc := range item.Exceptions {
		if err := exc.validate(feed); err != nil {
			return err
		}
	}

	return nil

}

func (item *Calendar) validate(feed *Feed) error {
	if item.ServiceId == "" {
		return &RequirementError{
			Struct: "CalendarSerializable",
			Field:  "ServiceId",
		}
	}

	if item.StartDate.IsEmpty() {
		return &RequirementError{
			Struct: "CalendarSerializable",
			Field:  "StartDate",
		}
	}

	if item.EndDate.IsEmpty() {
		return &RequirementError{
			Struct: "CalendarSerializable",
			Field:  "EndDate",
		}
	}

	return nil
}

func (item *CalendarSerializable) validate(feed *FeedSerializable) error {
	if item.ServiceId == "" {
		return &RequirementError{
			Struct: "CalendarSerializable",
			Field:  "ServiceId",
		}
	}

	if item.StartDate.IsEmpty() {
		return &RequirementError{
			Struct: "CalendarSerializable",
			Field:  "StartDate",
		}
	}

	if item.EndDate.IsEmpty() {
		return &RequirementError{
			Struct: "CalendarSerializable",
			Field:  "EndDate",
		}
	}

	return nil
}

func (item *CalendarDate) validate(feed *Feed) error {
	if item.ServiceId == "" {
		return &RequirementError{
			Struct: "CalendarDateSerializable",
			Field:  "ServiceId",
		}
	}

	if item.Date.IsEmpty() {
		return &RequirementError{
			Struct: "CalendarDateSerializable",
			Field:  "Date",
		}
	}

	if !in(item.Type, 1, 2) {
		return &InvalidValueError{
			Struct: "CalendarDateSerializable",
			Field:  "Type",
			Value:  item.Type,
		}
	}

	return nil
}

func (item *CalendarDateSerializable) validate(feed *FeedSerializable) error {

	if item.ServiceId == "" {
		return &RequirementError{
			Struct: "CalendarDateSerializable",
			Field:  "ServiceId",
		}
	}

	if item.Date.IsEmpty() {
		return &RequirementError{
			Struct: "CalendarDateSerializable",
			Field:  "Date",
		}
	}

	if !in(item.Type, 1, 2) {
		return &InvalidValueError{
			Struct: "CalendarDateSerializable",
			Field:  "Type",
			Value:  item.Type,
		}
	}

	return nil
}

func (item *FareAttribute) validate(feed *Feed) error {
	if item.Id == "" {
		return &RequirementError{
			Struct: "FareAttributeSerializable",
			Field:  "Id",
		}
	}

	if item.Price < 0 {
		return &InvalidValueError{
			Struct: "FareAttributeSerializable",
			Field:  "Price",
			Value:  item.Price,
		}
	}

	if item.CurrencyType == "" {
		return &RequirementError{
			Struct: "FareAttributeSerializable",
			Field:  "CurrencyType",
		}
	}

	if !in(item.PaymentMethod, 0, 1) {
		return &InvalidValueError{
			Struct: "FareAttributeSerializable",
			Field:  "PaymentMethod",
			Value:  item.PaymentMethod,
		}
	}

	if !in(item.Transfers, -1, 0, 1, 2) {
		return &InvalidValueError{
			Struct: "FareAttributeSerializable",
			Field:  "Transfers",
			Value:  item.Transfers,
		}
	}

	if len(feed.Agencies) > 1 && item.Agency == nil {
		return &ConditionnalRequirementError{
			Struct:    "FareAttributeSerializable",
			Field:     "AgencyId",
			Condition: "This field is required for datasets with multiple agencies, otherwise it is optional.",
		}
	}

	if item.Agency != nil {
		if _, ok := feed.Agencies[item.Agency.Id]; !ok {
			return &ReferenceError{
				Struct:      "FareAttributeSerializable",
				Field:       "AgencyId",
				Destination: "AgencySerializable",
				Value:       item.Agency.Id,
			}
		}
	}

	if item.TransferDuration < 0 {
		return &InvalidValueError{
			Struct: "FareAttributeSerializable",
			Field:  "TransferDuration",
			Value:  item.TransferDuration,
		}
	}

	return nil
}

func (item *FareAttributeSerializable) validate(feed *FeedSerializable) error {

	if item.Id == "" {
		return &RequirementError{
			Struct: "FareAttributeSerializable",
			Field:  "Id",
		}
	}

	if item.Price < 0 {
		return &InvalidValueError{
			Struct: "FareAttributeSerializable",
			Field:  "Price",
			Value:  item.Price,
		}
	}

	if item.CurrencyType == "" {
		return &RequirementError{
			Struct: "FareAttributeSerializable",
			Field:  "CurrencyType",
		}
	}

	if !in(item.PaymentMethod, 0, 1) {
		return &InvalidValueError{
			Struct: "FareAttributeSerializable",
			Field:  "PaymentMethod",
			Value:  item.PaymentMethod,
		}
	}

	if !in(item.Transfers, -1, 0, 1, 2) {
		return &InvalidValueError{
			Struct: "FareAttributeSerializable",
			Field:  "Transfers",
			Value:  item.Transfers,
		}
	}

	if len(feed.Agencies) > 1 && item.AgencyId == "" {
		return &ConditionnalRequirementError{
			Struct:    "FareAttributeSerializable",
			Field:     "AgencyId",
			Condition: "This field is required for datasets with multiple agencies, otherwise it is optional.",
		}
	}

	if item.AgencyId != "" {
		if _, ok := feed.Agencies[item.AgencyId]; !ok {
			return &ReferenceError{
				Struct:      "FareAttributeSerializable",
				Field:       "AgencyId",
				Destination: "AgencySerializable",
				Value:       item.AgencyId,
			}
		}
	}

	if item.TransferDuration < 0 {
		return &InvalidValueError{
			Struct: "FareAttributeSerializable",
			Field:  "TransferDuration",
			Value:  item.TransferDuration,
		}
	}

	return nil
}

func (item *FareRule) validate(feed *Feed) error {
	if item.FareId == "" {
		return &RequirementError{
			Struct: "FareRuleSerializable",
			Field:  "FareId",
		}
	} else {
		if _, ok := feed.FareAttributes[item.FareId]; !ok {
			return &ReferenceError{
				Struct:      "FareRuleSerializable",
				Field:       "FareId",
				Destination: "FareAttributeSerializable",
				Value:       item.FareId,
			}
		}
	}

	if item.Route != nil {
		if _, ok := feed.Routes[item.Route.Id]; !ok {
			return &ReferenceError{
				Struct:      "FareRuleSerializable",
				Field:       "RouteId",
				Destination: "RouteSerializable",
				Value:       item.Route.Id,
			}
		}
	}

	searchingOrigin := item.Origin != nil
	searchingDestination := item.Destination != nil
	searchingContains := item.Contains != nil

	for _, stop := range feed.Stops {
		if searchingOrigin {
			if stop.Zone.Id == item.Origin.Id {
				searchingOrigin = false
			}
		}

		if searchingDestination {
			if stop.Zone.Id == item.Destination.Id {
				searchingDestination = false
			}
		}

		if searchingContains {
			if stop.Zone.Id == item.Contains.Id {
				searchingContains = false
			}
		}

		if !searchingOrigin && !searchingDestination && !searchingContains {
			break
		}
	}

	if searchingOrigin {
		return &ReferenceError{
			Struct:      "FareRuleSerializable",
			Field:       "OriginId",
			Destination: "StopSerializable",
			Value:       item.Origin.Id,
		}
	}

	if searchingDestination {
		return &ReferenceError{
			Struct:      "FareRuleSerializable",
			Field:       "DestinationId",
			Destination: "StopSerializable",
			Value:       item.Destination.Id,
		}
	}

	if searchingContains {
		return &ReferenceError{
			Struct:      "FareRuleSerializable",
			Field:       "ContainsId",
			Destination: "StopSerializable",
			Value:       item.Contains.Id,
		}
	}

	return nil
}

func (item *FareRuleSerializable) validate(feed *FeedSerializable) error {

	if item.FareId == "" {
		return &RequirementError{
			Struct: "FareRuleSerializable",
			Field:  "FareId",
		}
	} else {
		if _, ok := feed.FareAttributes[item.FareId]; !ok {
			return &ReferenceError{
				Struct:      "FareRuleSerializable",
				Field:       "FareId",
				Destination: "FareAttributeSerializable",
				Value:       item.FareId,
			}
		}
	}

	if item.RouteId != "" {
		if _, ok := feed.Routes[item.RouteId]; !ok {
			return &ReferenceError{
				Struct:      "FareRuleSerializable",
				Field:       "RouteId",
				Destination: "RouteSerializable",
				Value:       item.RouteId,
			}
		}
	}

	searchingOrigin := item.OriginId != ""
	searchingDestination := item.DestinationId != ""
	searchingContains := item.ContainsId != ""

	for _, stop := range feed.Stops {
		if searchingOrigin {
			if stop.ZoneId == item.OriginId {
				searchingOrigin = false
			}
		}

		if searchingDestination {
			if stop.ZoneId == item.DestinationId {
				searchingDestination = false
			}
		}

		if searchingContains {
			if stop.ZoneId == item.ContainsId {
				searchingContains = false
			}
		}

		if !searchingOrigin && !searchingDestination && !searchingContains {
			break
		}
	}

	if searchingOrigin {
		return &ReferenceError{
			Struct:      "FareRuleSerializable",
			Field:       "OriginId",
			Destination: "StopSerializable",
			Value:       item.OriginId,
		}
	}

	if searchingDestination {
		return &ReferenceError{
			Struct:      "FareRuleSerializable",
			Field:       "DestinationId",
			Destination: "StopSerializable",
			Value:       item.DestinationId,
		}
	}

	if searchingContains {
		return &ReferenceError{
			Struct:      "FareRuleSerializable",
			Field:       "ContainsId",
			Destination: "StopSerializable",
			Value:       item.ContainsId,
		}
	}

	return nil
}

func (item *Shape) validate(feed *Feed) error {
	if item.Id == "" {
		return &RequirementError{
			Struct: "ShapeSerializable",
			Field:  "ShapeId",
		}
	}

	for _, pt := range item.Points {
		if pt.ShapeId != item.Id {
			return &ReferenceError{
				Struct:      "ShapePoint",
				Field:       "ShapeId",
				Destination: "Shape",
				Value:       pt.ShapeId,
			}
		}

		if pt.Lat == 0 {
			return &RequirementError{
				Struct: "ShapeSerializable",
				Field:  "Lat",
			}
		}

		if pt.Long == 0 {
			return &RequirementError{
				Struct: "ShapeSerializable",
				Field:  "Long",
			}
		}

		if pt.Sequence < 0 {
			return &InvalidValueError{
				Struct: "ShapeSerializable",
				Field:  "Sequence",
				Value:  pt.Sequence,
			}
		}

		if pt.DistTraveled < 0 {
			return &InvalidValueError{
				Struct: "ShapeSerializable",
				Field:  "DistTraveled",
				Value:  pt.DistTraveled,
			}
		}
	}

	return nil
}

func (item *ShapeSerializable) validate(feed *FeedSerializable) error {
	if item.ShapeId == "" {
		return &RequirementError{
			Struct: "ShapeSerializable",
			Field:  "ShapeId",
		}
	}

	if item.Lat == 0 {
		return &RequirementError{
			Struct: "ShapeSerializable",
			Field:  "Lat",
		}
	}

	if item.Long == 0 {
		return &RequirementError{
			Struct: "ShapeSerializable",
			Field:  "Long",
		}
	}

	if item.Sequence < 0 {
		return &InvalidValueError{
			Struct: "ShapeSerializable",
			Field:  "Sequence",
			Value:  item.Sequence,
		}
	}

	if item.DistTraveled < 0 {
		return &InvalidValueError{
			Struct: "ShapeSerializable",
			Field:  "DistTraveled",
			Value:  item.DistTraveled,
		}
	}

	return nil
}

func (item *Frequency) validate(feed *Feed) error {
	if item.Trip == nil {
		return &RequirementError{
			Struct: "FrequencySerializable",
			Field:  "TripId",
		}
	} else {
		if _, ok := feed.Trips[item.Trip.Id]; !ok {
			return &ReferenceError{
				Struct:      "FrequencySerializable",
				Field:       "TripId",
				Destination: "TripSerializable",
				Value:       item.Trip.Id,
			}
		}
	}

	if item.StartTime.IsEmpty() {
		return &RequirementError{
			Struct: "FrequencySerializable",
			Field:  "StartTime",
		}
	}

	if item.EndTime.IsEmpty() {
		return &RequirementError{
			Struct: "FrequencySerializable",
			Field:  "EndTime",
		}
	}

	if item.HeadwaySecs < 0 {
		return &InvalidValueError{
			Struct: "FrequencySerializable",
			Field:  "HeadwaySecs",
			Value:  item.HeadwaySecs,
		}
	}

	if !in(item.ExactTimes, 0, 1) {
		return &InvalidValueError{
			Struct: "FrequencySerializable",
			Field:  "ExactTimes",
			Value:  item.ExactTimes,
		}
	}

	return nil
}

func (item *FrequencySerializable) validate(feed *FeedSerializable) error {

	if item.TripId == "" {
		return &RequirementError{
			Struct: "FrequencySerializable",
			Field:  "TripId",
		}
	} else {
		if _, ok := feed.Trips[item.TripId]; !ok {
			return &ReferenceError{
				Struct:      "FrequencySerializable",
				Field:       "TripId",
				Destination: "TripSerializable",
				Value:       item.TripId,
			}
		}
	}

	if item.StartTime.IsEmpty() {
		return &RequirementError{
			Struct: "FrequencySerializable",
			Field:  "StartTime",
		}
	}

	if item.EndTime.IsEmpty() {
		return &RequirementError{
			Struct: "FrequencySerializable",
			Field:  "EndTime",
		}
	}

	if item.HeadwaySecs < 0 {
		return &InvalidValueError{
			Struct: "FrequencySerializable",
			Field:  "HeadwaySecs",
			Value:  item.HeadwaySecs,
		}
	}

	if !in(item.ExactTimes, 0, 1) {
		return &InvalidValueError{
			Struct: "FrequencySerializable",
			Field:  "ExactTimes",
			Value:  item.ExactTimes,
		}
	}

	return nil
}

func (item *Transfer) validate(feed *Feed) error {
	if item.From == nil {
		return &RequirementError{
			Struct: "TransferSerializable",
			Field:  "From",
		}
	} else {
		if _, ok := feed.Stops[item.From.Id]; !ok {
			return &ReferenceError{
				Struct:      "TransferSerializable",
				Field:       "From",
				Destination: "StopSerializable",
				Value:       item.From.Id,
			}
		}
	}

	if item.To == nil {
		return &RequirementError{
			Struct: "TransferSerializable",
			Field:  "To",
		}
	} else {
		if _, ok := feed.Stops[item.To.Id]; !ok {
			return &ReferenceError{
				Struct:      "TransferSerializable",
				Field:       "To",
				Destination: "StopSerializable",
				Value:       item.To.Id,
			}
		}
	}

	if !in(item.Type, 0, 1, 2, 3) {
		return &InvalidValueError{
			Struct: "TransferSerializable",
			Field:  "Type",
			Value:  item.Type,
		}
	}

	if item.MinTransferTime < 0 {
		return &InvalidValueError{
			Struct: "TransferSerializable",
			Field:  "MinTransferTime",
			Value:  item.MinTransferTime,
		}
	}

	return nil
}

func (item *TransferSerializable) validate(feed *FeedSerializable) error {

	if item.From == "" {
		return &RequirementError{
			Struct: "TransferSerializable",
			Field:  "From",
		}
	} else {
		if _, ok := feed.Stops[item.From]; !ok {
			return &ReferenceError{
				Struct:      "TransferSerializable",
				Field:       "From",
				Destination: "StopSerializable",
				Value:       item.From,
			}
		}
	}

	if item.To == "" {
		return &RequirementError{
			Struct: "TransferSerializable",
			Field:  "To",
		}
	} else {
		if _, ok := feed.Stops[item.To]; !ok {
			return &ReferenceError{
				Struct:      "TransferSerializable",
				Field:       "To",
				Destination: "StopSerializable",
				Value:       item.To,
			}
		}
	}

	if !in(item.Type, 0, 1, 2, 3) {
		return &InvalidValueError{
			Struct: "TransferSerializable",
			Field:  "Type",
			Value:  item.Type,
		}
	}

	if item.MinTransferTime < 0 {
		return &InvalidValueError{
			Struct: "TransferSerializable",
			Field:  "MinTransferTime",
			Value:  item.MinTransferTime,
		}
	}

	return nil
}

func (item *Pathway) validate(feed *Feed) error {
	if item.Id == "" {
		return &RequirementError{
			Struct: "PathwaySerializable",
			Field:  "Id",
		}
	}

	if item.From == nil {
		return &RequirementError{
			Struct: "PathwaySerializable",
			Field:  "From",
		}
	} else {
		if _, ok := feed.Stops[item.From.Id]; !ok {
			return &ReferenceError{
				Struct:      "PathwaySerializable",
				Field:       "From",
				Destination: "StopSerializable",
				Value:       item.From.Id,
			}
		}
	}

	if item.To == nil {
		return &RequirementError{
			Struct: "PathwaySerializable",
			Field:  "To",
		}
	} else {
		if _, ok := feed.Stops[item.To.Id]; !ok {
			return &ReferenceError{
				Struct:      "PathwaySerializable",
				Field:       "To",
				Destination: "StopSerializable",
				Value:       item.To.Id,
			}
		}
	}

	if !in(item.Mode, 1, 2, 3, 4, 5, 6, 7) {
		return &InvalidValueError{
			Struct: "PathwaySerializable",
			Field:  "Mode",
			Value:  item.Mode,
		}
	}

	if !in(item.IsBidirectional, 0, 1) {
		return &InvalidValueError{
			Struct: "PathwaySerializable",
			Field:  "IsBidirectional",
			Value:  item.IsBidirectional,
		}
	}

	if item.Length < 0 {
		return &InvalidValueError{
			Struct: "PathwaySerializable",
			Field:  "Length",
			Value:  item.Length,
		}
	}

	if item.TraversalTime < 0 {
		return &InvalidValueError{
			Struct: "PathwaySerializable",
			Field:  "TraversalTime",
			Value:  item.TraversalTime,
		}
	}

	if item.MinWidth < 0 {
		return &InvalidValueError{
			Struct: "PathwaySerializable",
			Field:  "MinWidth",
			Value:  item.MinWidth,
		}
	}

	return nil
}

func (item *PathwaySerializable) validate(feed *FeedSerializable) error {

	if item.Id == "" {
		return &RequirementError{
			Struct: "PathwaySerializable",
			Field:  "Id",
		}
	}

	if item.From == "" {
		return &RequirementError{
			Struct: "PathwaySerializable",
			Field:  "From",
		}
	} else {
		if _, ok := feed.Stops[item.From]; !ok {
			return &ReferenceError{
				Struct:      "PathwaySerializable",
				Field:       "From",
				Destination: "StopSerializable",
				Value:       item.From,
			}
		}
	}

	if item.To == "" {
		return &RequirementError{
			Struct: "PathwaySerializable",
			Field:  "To",
		}
	} else {
		if _, ok := feed.Stops[item.To]; !ok {
			return &ReferenceError{
				Struct:      "PathwaySerializable",
				Field:       "To",
				Destination: "StopSerializable",
				Value:       item.To,
			}
		}
	}

	if !in(item.Mode, 1, 2, 3, 4, 5, 6, 7) {
		return &InvalidValueError{
			Struct: "PathwaySerializable",
			Field:  "Mode",
			Value:  item.Mode,
		}
	}

	if !in(item.IsBidirectional, 0, 1) {
		return &InvalidValueError{
			Struct: "PathwaySerializable",
			Field:  "IsBidirectional",
			Value:  item.IsBidirectional,
		}
	}

	if item.Length < 0 {
		return &InvalidValueError{
			Struct: "PathwaySerializable",
			Field:  "Length",
			Value:  item.Length,
		}
	}

	if item.TraversalTime < 0 {
		return &InvalidValueError{
			Struct: "PathwaySerializable",
			Field:  "TraversalTime",
			Value:  item.TraversalTime,
		}
	}

	if item.MinWidth < 0 {
		return &InvalidValueError{
			Struct: "PathwaySerializable",
			Field:  "MinWidth",
			Value:  item.MinWidth,
		}
	}

	return nil
}

func (item *Level) validate(feed *Feed) error {
	if item.Id == "" {
		return &RequirementError{
			Struct: "LevelSerializable",
			Field:  "Id",
		}
	}
	return nil
}

func (item *LevelSerializable) validate(feed *FeedSerializable) error {
	if item.Id == "" {
		return &RequirementError{
			Struct: "LevelSerializable",
			Field:  "Id",
		}
	}
	return nil
}

func (item *FeedInfo) validate(feed *Feed) error {

	if item.PublisherName == "" {
		return &RequirementError{
			Struct: "FeedInfo",
			Field:  "PublisherName",
		}
	}

	if item.PublisherUrl == "" {
		return &RequirementError{
			Struct: "FeedInfo",
			Field:  "PublisherUrl",
		}
	}

	if item.FeedLang == "" {
		return &RequirementError{
			Struct: "FeedInfo",
			Field:  "FeedLang",
		}
	}

	return nil
}

func (item *FeedInfoSerializable) validate(feed *FeedSerializable) error {

	if item.PublisherName == "" {
		return &RequirementError{
			Struct: "FeedInfoSerializable",
			Field:  "PublisherName",
		}
	}

	if item.PublisherUrl == "" {
		return &RequirementError{
			Struct: "FeedInfoSerializable",
			Field:  "PublisherUrl",
		}
	}

	if item.FeedLang == "" {
		return &RequirementError{
			Struct: "FeedInfoSerializable",
			Field:  "FeedLang",
		}
	}

	return nil
}

func (item *Translation) validate(feed *Feed) error {

	if !in(item.TableName, "agency", "stops", "routes", "trips", "stop_times",
		"feed_info", "pathways", "levels", "attributions") {
		return &InvalidValueError{
			Struct: "Translation",
			Field:  "TableName",
			Value:  item.TableName,
		}
	}

	if item.FieldName == "" {
		return &RequirementError{
			Struct: "Translation",
			Field:  "FieldName",
		}
	}

	if item.Language == "" {
		return &RequirementError{
			Struct: "Translation",
			Field:  "Language",
		}
	}

	if item.Translation == "" {
		return &RequirementError{
			Struct: "Translation",
			Field:  "Translation",
		}
	}

	if item.TableName == "feed_info" && item.RecordId != "" {
		return &ConditionnalRequirementError{
			Struct:    "Translation",
			Field:     "RecordId",
			Condition: "Forbidden if TableName equals feed_info.",
		}
	}

	if item.FieldValue != "" && item.RecordId != "" {
		return &ConditionnalRequirementError{
			Struct:    "Translation",
			Field:     "RecordId",
			Condition: "Forbidden if FieldValue is defined.",
		}
	}

	if item.FieldValue == "" && item.RecordId == "" {
		return &ConditionnalRequirementError{
			Struct:    "Translation",
			Field:     "RecordId",
			Condition: "Required if field_value is empty.",
		}
	}

	if item.TableName == "feed_info" && item.RecordSubId != "" {
		return &ConditionnalRequirementError{
			Struct:    "Translation",
			Field:     "RecordSubId",
			Condition: "Forbidden if TableName equals feed_info.",
		}
	}

	if item.FieldValue != "" && item.RecordSubId != "" {
		return &ConditionnalRequirementError{
			Struct:    "Translation",
			Field:     "RecordSubId",
			Condition: "Forbidden if FieldValue is defined.",
		}
	}

	if item.TableName == "stop_times" && item.RecordId != "" && item.RecordSubId == "" {
		return &ConditionnalRequirementError{
			Struct:    "Translation",
			Field:     "RecordSubId",
			Condition: "Required if TableName equals stop_times and RecorId is defined.",
		}
	}

	if item.TableName == "feed_info" && item.FieldValue != "" {
		return &ConditionnalRequirementError{
			Struct:    "Translation",
			Field:     "FieldValue",
			Condition: "Forbidden if TableName equals feed_info.",
		}
	}

	if item.RecordId != "" && item.FieldValue != "" {
		return &ConditionnalRequirementError{
			Struct:    "Translation",
			Field:     "FieldValue",
			Condition: "Forbidden if RecordId is defined.",
		}
	}

	if item.RecordId == "" && item.FieldValue == "" {
		return &ConditionnalRequirementError{
			Struct:    "Translation",
			Field:     "FieldValue",
			Condition: "Required if RecordId is empty.",
		}
	}

	return nil
}

func (item *TranslationSerializable) validate(feed *FeedSerializable) error {

	if !in(item.TableName, "agency", "stops", "routes", "trips", "stop_times",
		"feed_info", "pathways", "levels", "attributions") {
		return &InvalidValueError{
			Struct: "TranslationSerializable",
			Field:  "TableName",
			Value:  item.TableName,
		}
	}

	if item.FieldName == "" {
		return &RequirementError{
			Struct: "TranslationSerializable",
			Field:  "FieldName",
		}
	}

	if item.Language == "" {
		return &RequirementError{
			Struct: "TranslationSerializable",
			Field:  "Language",
		}
	}

	if item.Translation == "" {
		return &RequirementError{
			Struct: "TranslationSerializable",
			Field:  "Translation",
		}
	}

	if item.TableName == "feed_info" && item.RecordId != "" {
		return &ConditionnalRequirementError{
			Struct:    "TranslationSerializable",
			Field:     "RecordId",
			Condition: "Forbidden if TableName equals feed_info.",
		}
	}

	if item.FieldValue != "" && item.RecordId != "" {
		return &ConditionnalRequirementError{
			Struct:    "TranslationSerializable",
			Field:     "RecordId",
			Condition: "Forbidden if FieldValue is defined.",
		}
	}

	if item.FieldValue == "" && item.RecordId == "" {
		return &ConditionnalRequirementError{
			Struct:    "TranslationSerializable",
			Field:     "RecordId",
			Condition: "Required if field_value is empty.",
		}
	}

	if item.TableName == "feed_info" && item.RecordSubId != "" {
		return &ConditionnalRequirementError{
			Struct:    "TranslationSerializable",
			Field:     "RecordSubId",
			Condition: "Forbidden if TableName equals feed_info.",
		}
	}

	if item.FieldValue != "" && item.RecordSubId != "" {
		return &ConditionnalRequirementError{
			Struct:    "TranslationSerializable",
			Field:     "RecordSubId",
			Condition: "Forbidden if FieldValue is defined.",
		}
	}

	if item.TableName == "stop_times" && item.RecordId != "" && item.RecordSubId == "" {
		return &ConditionnalRequirementError{
			Struct:    "TranslationSerializable",
			Field:     "RecordSubId",
			Condition: "Required if TableName equals stop_times and RecorId is defined.",
		}
	}

	if item.TableName == "feed_info" && item.FieldValue != "" {
		return &ConditionnalRequirementError{
			Struct:    "TranslationSerializable",
			Field:     "FieldValue",
			Condition: "Forbidden if TableName equals feed_info.",
		}
	}

	if item.RecordId != "" && item.FieldValue != "" {
		return &ConditionnalRequirementError{
			Struct:    "TranslationSerializable",
			Field:     "FieldValue",
			Condition: "Forbidden if RecordId is defined.",
		}
	}

	if item.RecordId == "" && item.FieldValue == "" {
		return &ConditionnalRequirementError{
			Struct:    "TranslationSerializable",
			Field:     "FieldValue",
			Condition: "Required if RecordId is empty.",
		}
	}

	return nil
}

func (item *Attribution) validate(feed *Feed) error {

	if item.Agency != nil {
		if _, ok := feed.Agencies[item.Agency.Id]; !ok {
			return &ReferenceError{
				Struct:      "Attribution",
				Field:       "Agency",
				Destination: "Agency",
				Value:       item.Agency.Id,
			}
		}
	}

	if item.Route != nil {
		if _, ok := feed.Routes[item.Route.Id]; !ok {
			return &ReferenceError{
				Struct:      "Attribution",
				Field:       "Route",
				Destination: "Route",
				Value:       item.Route.Id,
			}
		}
	}

	if item.Trip != nil {
		if _, ok := feed.Routes[item.Trip.Id]; !ok {
			return &ReferenceError{
				Struct:      "Attribution",
				Field:       "Trip",
				Destination: "Trip",
				Value:       item.Trip.Id,
			}
		}
	}

	if item.OrganizationName == "" {
		return &RequirementError{
			Struct: "Attribution",
			Field:  "OrganizationName",
		}
	}

	if item.IsProducer != 1 && item.IsOperator != 1 && item.IsAuthority != 1 {
		return &ConditionnalRequirementError{
			Struct:    "Attribution",
			Field:     "IsProducer",
			Condition: "At least one of the fields, either IsProducer, IsOperator, or IsAuthority, must be set at 1.",
		}
	}

	return nil
}

func (item *AttributionSerializable) validate(feed *FeedSerializable) error {

	if item.AgencyId != "" {
		if _, ok := feed.Agencies[item.AgencyId]; !ok {
			return &ReferenceError{
				Struct:      "AttributionSerializable",
				Field:       "AgencyId",
				Destination: "AgencySerializable",
				Value:       item.AgencyId,
			}
		}
	}

	if item.RouteId != "" {
		if _, ok := feed.Routes[item.RouteId]; !ok {
			return &ReferenceError{
				Struct:      "AttributionSerializable",
				Field:       "RouteId",
				Destination: "RouteSerializable",
				Value:       item.RouteId,
			}
		}
	}

	if item.TripId != "" {
		if _, ok := feed.Routes[item.TripId]; !ok {
			return &ReferenceError{
				Struct:      "AttributionSerializable",
				Field:       "TripId",
				Destination: "TripSerializable",
				Value:       item.TripId,
			}
		}
	}

	if item.OrganizationName == "" {
		return &RequirementError{
			Struct: "AttributionSerializable",
			Field:  "OrganizationName",
		}
	}

	if item.IsProducer != 1 && item.IsOperator != 1 && item.IsAuthority != 1 {
		return &ConditionnalRequirementError{
			Struct:    "AttributionSerializable",
			Field:     "IsProducer",
			Condition: "At least one of the fields, either IsProducer, IsOperator, or IsAuthority, must be set at 1.",
		}
	}

	return nil
}

func in(item interface{}, array ...interface{}) bool {
	for _, elem := range array {
		if item == elem {
			return true
		}
	}
	return false
}
