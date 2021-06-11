package gtfs

type linker struct{}

func (l linker) Link(feed *FeedSerializable) (*Feed, error) {

	f := &Feed{
		Agencies:       make(map[string]*Agency, len(feed.Agencies)),
		Levels:         make(map[string]*Level, len(feed.Levels)),
		Zones:          make(map[string]*Zone),
		Stops:          make(map[string]*Stop, len(feed.Stops)),
		Transfers:      make([]*Transfer, 0, len(feed.Transfers)),
		Pathways:       make(map[string]*Pathway, len(feed.Pathways)),
		Services:       make(map[string]*Service, len(feed.Calendar)),
		Shapes:         make(map[string]*Shape),
		Routes:         make(map[string]*Route, len(feed.Routes)),
		Trips:          make(map[string]*Trip, len(feed.Trips)),
		Block:          make(map[string]*Block),
		Frequencies:    make(map[string]*Frequency, len(feed.Frequencies)),
		FareAttributes: make(map[string]*FareAttribute, len(feed.FareAttributes)),
		Translations:   make([]*Translation, len(feed.Translations)),
		Attributions:   make([]*Attribution, len(feed.Attributions)),
	}

	// TODO: might benefit from goroutine by creating different group to run in
	// parallel while keeping the dependency order
	err := l.Agencies(f, feed)
	if err == nil {
		err = l.Levels(f, feed)
	}
	if err == nil {
		err = l.Stops(f, feed)
	}
	if err == nil {
		err = l.Transfers(f, feed)
	}
	if err == nil {
		err = l.Pathways(f, feed)
	}
	if err == nil {
		err = l.Services(f, feed)
	}
	if err == nil {
		err = l.Shapes(f, feed)
	}
	if err == nil {
		err = l.Routes(f, feed)
	}
	if err == nil {
		err = l.Trips(f, feed)
	}
	if err == nil {
		err = l.StopTimes(f, feed)
	}
	if err == nil {
		err = l.Frequencies(f, feed)
	}
	if err == nil {
		err = l.Fares(f, feed)
	}
	if err != nil {
		err = l.FeedInfo(f, feed)
	}
	if err != nil {
		err = l.Translations(f, feed)
	}
	if err != nil {
		err = l.Attributions(f, feed)
	}

	if err != nil {
		return nil, err
	}

	return f, err
}

func (l linker) Agencies(f *Feed, feed *FeedSerializable) error {
	for _, item := range feed.Agencies {
		f.Agencies[item.Id] = &Agency{
			Id:       item.Id,
			Name:     item.Name,
			Url:      item.Url,
			Timezone: item.Timezone,
			Lang:     item.Lang,
			Phone:    item.Phone,
			FareUrl:  item.FareUrl,
			Email:    item.Email,
		}
	}
	return nil
}

func (l linker) Levels(f *Feed, feed *FeedSerializable) error {
	for _, item := range feed.Levels {
		f.Levels[item.Id] = &Level{
			Id:    item.Id,
			Index: item.Index,
			Name:  item.Name,
		}
	}
	return nil
}

func (l linker) Stops(f *Feed, feed *FeedSerializable) error {
	for _, item := range feed.Stops {

		zone, ok := f.Zones[item.ZoneId]
		if !ok {
			zone = &Zone{Id: item.ZoneId}
			f.Zones[zone.Id] = zone
		}

		f.Stops[item.Id] = &Stop{
			Id:                 item.Id,
			Code:               item.Code,
			Name:               item.Name,
			Desc:               item.Desc,
			Lat:                item.Lat,
			Long:               item.Long,
			Url:                item.Url,
			LocationType:       item.LocationType,
			Timezone:           item.Timezone,
			WheelchairBoarding: item.WheelchairBoarding,
			PlatformCode:       item.PlatformCode,
			Level:              f.Levels[item.LevelId],
			Zone:               zone,
			TransfersFrom:      make(map[string]*Transfer),
			TransfersTo:        make(map[string]*Transfer),
		}
	}

	for _, item := range f.Stops {
		id := feed.Stops[item.Id].ParentStation
		item.ParentStation = f.Stops[id]
	}

	return nil
}

func (l linker) Transfers(f *Feed, feed *FeedSerializable) error {
	for _, item := range feed.Transfers {
		f.Transfers = append(f.Transfers, &Transfer{
			From:            f.Stops[item.From],
			To:              f.Stops[item.To],
			Type:            item.Type,
			MinTransferTime: item.MinTransferTime,
		})
	}
	return nil
}

func (l linker) Pathways(f *Feed, feed *FeedSerializable) error {
	for _, item := range feed.Pathways {
		f.Pathways[item.Id] = &Pathway{
			Id:                   item.Id,
			From:                 f.Stops[item.From],
			To:                   f.Stops[item.To],
			Mode:                 item.Mode,
			IsBidirectional:      item.IsBidirectional,
			Length:               item.Length,
			TraversalTime:        item.TraversalTime,
			StairCount:           item.StairCount,
			MaxSlope:             item.MaxSlope,
			MinWidth:             item.MinWidth,
			SignpostedAs:         item.SignpostedAs,
			ReversedSignpostedAs: item.ReversedSignpostedAs,
		}
	}
	return nil
}

func (l linker) Services(f *Feed, feed *FeedSerializable) error {
	for _, item := range feed.Calendar {
		service, ok := f.Services[item.ServiceId]
		if !ok {
			service = &Service{Id: item.ServiceId}
			f.Services[service.Id] = service
		}

		service.Calendar = &Calendar{
			ServiceId: item.ServiceId,
			Days:      item.Days, //TODO: might need deep copy i don't know
			StartDate: item.StartDate,
			EndDate:   item.EndDate,
		}
	}

	for id, serv := range feed.CalendarDates {
		service, ok := f.Services[id]
		if !ok {
			service = &Service{Id: id}
			f.Services[service.Id] = service
		}

		for _, item := range serv {
			service.Exceptions = append(service.Exceptions, &CalendarDate{
				ServiceId: item.ServiceId,
				Date:      item.Date,
				Type:      item.Type,
			})
		}
	}

	return nil
}

func (l linker) Shapes(f *Feed, feed *FeedSerializable) error {
	for id, shp := range feed.Shapes {
		shape, ok := f.Shapes[id]
		if !ok {
			shape = &Shape{Id: id}
			f.Shapes[id] = shape
		}

		for _, item := range shp {
			shape.Points = append(shape.Points, &ShapePoint{
				ShapeId:      item.ShapeId,
				Lat:          item.Lat,
				Long:         item.Long,
				Sequence:     item.Sequence,
				DistTraveled: item.DistTraveled,
			})
		}

		// TODO: Sort the points in the sequence order
	}
	return nil
}

func (l linker) Routes(f *Feed, feed *FeedSerializable) error {
	for _, item := range feed.Routes {
		f.Routes[item.Id] = &Route{
			Id:                item.Id,
			Agency:            f.Agencies[item.AgencyId],
			ShortName:         item.ShortName,
			LongName:          item.LongName,
			Desc:              item.Desc,
			Type:              item.Type,
			Url:               item.Url,
			Color:             item.Color,
			TextColor:         item.TextColor,
			SortOrder:         item.SortOrder,
			ContinuousPickup:  item.ContinuousPickup,
			ContinuousDropOff: item.ContinuousDropOff,
		}
	}
	return nil
}

func (l linker) Trips(f *Feed, feed *FeedSerializable) error {
	for _, item := range feed.Trips {

		block, ok := f.Block[item.BlockId]
		if !ok {
			block = &Block{Id: item.BlockId}
			f.Block[block.Id] = block
		}

		route, ok := f.Routes[item.RouteId]
		if ok {
			trip := &Trip{
				Id:                   item.Id,
				Route:                route,
				Service:              f.Services[item.ServiceId],
				Block:                block,
				Shape:                f.Shapes[item.ShapeId],
				Headsign:             item.Headsign,
				ShortName:            item.ShortName,
				DirectionId:          item.DirectionId,
				WheelchairAccessible: item.WheelchairAccessible,
				BikesAllowed:         item.BikesAllowed,
			}

			route.Trips = append(route.Trips, trip)
			f.Trips[trip.Id] = trip
		}
	}
	return nil
}

func (l linker) StopTimes(f *Feed, feed *FeedSerializable) error {
	for id, array := range feed.StopTimes {
		trip, ok := f.Trips[id]
		if !ok {
			return &MissingStructError{
				Container: "Feed.Trips",
				Id:        id,
			}
		}

		for _, item := range array {
			trip.StopTimes = append(trip.StopTimes, &StopTime{
				Trip:              trip,
				Stop:              f.Stops[item.StopId],
				ArrivalTime:       item.ArrivalTime,
				DepartureTime:     item.DepartureTime,
				StopSequence:      item.StopSequence,
				StopHeadsign:      item.StopHeadsign,
				PickupType:        item.PickupType,
				DropOffType:       item.DropOffType,
				ContinuousPickup:  item.ContinuousPickup,
				ContinuousDropOff: item.ContinuousDropOff,
				ShapeDistTraveled: item.ShapeDistTraveled,
				Timepoint:         item.Timepoint,
			})
		}
	}
	return nil
}

func (l linker) Frequencies(f *Feed, feed *FeedSerializable) error {
	for _, item := range feed.Frequencies {
		f.Frequencies[item.TripId] = &Frequency{
			Trip:        f.Trips[item.TripId],
			StartTime:   item.StartTime,
			EndTime:     item.EndTime,
			HeadwaySecs: item.HeadwaySecs,
			ExactTimes:  item.ExactTimes,
		}
	}
	return nil
}

func (l linker) Fares(f *Feed, feed *FeedSerializable) error {
	for _, item := range feed.FareAttributes {
		f.FareAttributes[item.Id] = &FareAttribute{
			Id:               item.Id,
			Price:            item.Price,
			CurrencyType:     item.CurrencyType,
			PaymentMethod:    item.PaymentMethod,
			Transfers:        item.Transfers,
			Agency:           f.Agencies[item.AgencyId],
			TransferDuration: item.TransferDuration,
		}
	}

	for _, item := range feed.FareRules {
		fare, ok := f.FareAttributes[item.FareId]
		if !ok {
			return &MissingStructError{
				Container: "Feed.FareAttributes",
				Id:        item.FareId,
			}
		}

		fare.Rules = append(fare.Rules, &FareRule{
			FareId:      item.FareId,
			Route:       f.Routes[item.RouteId],
			Origin:      f.Zones[item.OriginId],
			Destination: f.Zones[item.DestinationId],
			Contains:    f.Zones[item.ContainsId],
		})
	}
	return nil
}

func (l linker) FeedInfo(f *Feed, feed *FeedSerializable) error {

	f.FeedInfo = &FeedInfo{
		PublisherName: feed.FeedInfo.PublisherName,
		PublisherUrl:  feed.FeedInfo.PublisherUrl,
		FeedLang:      feed.FeedInfo.FeedLang,
		DefaultLang:   feed.FeedInfo.DefaultLang,
		StartDate:     feed.FeedInfo.StartDate,
		EndDate:       feed.FeedInfo.EndDate,
		Version:       feed.FeedInfo.Version,
		ContactEmail:  feed.FeedInfo.ContactEmail,
		ContactUrl:    feed.FeedInfo.ContactUrl,
	}

	return nil
}

func (l linker) Translations(f *Feed, feed *FeedSerializable) error {

	for _, item := range feed.Translations {
		f.Translations = append(f.Translations, &Translation{
			TableName:   item.TableName,
			FieldName:   item.FieldName,
			Language:    item.Language,
			Translation: item.Translation,
			RecordId:    item.RecordId,
			RecordSubId: item.RecordSubId,
			FieldValue:  item.FieldValue,
		})
	}

	return nil
}

func (l linker) Attributions(f *Feed, feed *FeedSerializable) error {

	for _, item := range feed.Attributions {
		f.Attributions = append(f.Attributions, &Attribution{
			Id:               item.Id,
			Agency:           f.Agencies[item.AgencyId],
			Route:            f.Routes[item.RouteId],
			Trip:             f.Trips[item.TripId],
			OrganizationName: item.OrganizationName,
			IsProducer:       item.IsProducer,
			IsOperator:       item.IsOperator,
			IsAuthority:      item.IsAuthority,
			AttributionUrl:   item.AttributionUrl,
			AttributionEmail: item.AttributionEmail,
			AttributionPhone: item.AttributionPhone,
		})
	}

	return nil
}

type unlinker struct{}

func (l unlinker) Unlink(feed *Feed) (*FeedSerializable, error) {
	return nil, nil
}
