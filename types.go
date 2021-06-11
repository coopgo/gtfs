package gtfs

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// This file contains general purposes type for our gtfs library and extractor /
// convertor from string to these types.

// A general date type for the Date gtfs format. Since time within a service day
// can be above 24:00:00, a service day often contains information for the
// subsequent day(s).
//
// The empty value is created using Year == -1
//
// Example: 20180913 for September 13th, 2018.
type Date struct {
	Year  int
	Month int
	Day   int
}

var DateEmpty Date = Date{Year: -1}

func (d Date) IsEmpty() bool {
	return d == DateEmpty
}

// Parse the date from a string in gtfs format: YYYYMMDD
func ParseDate(str string) (Date, error) {
	if len(str) != 8 {
		return Date{}, fmt.Errorf("gtfs: invalid date format, %s", str)
	}

	var err error
	var date Date

	date.Year, err = strconv.Atoi(str[0:4])
	if err != nil || date.Year < 0 {
		return Date{}, fmt.Errorf("gtfs: invalid date format, %s", str)
	}

	date.Month, err = strconv.Atoi(str[4:6])
	if err != nil || date.Month < 0 || date.Month > 12 {
		return Date{}, fmt.Errorf("gtfs: invalid date format, %s", str)
	}

	date.Day, err = strconv.Atoi(str[6:8])
	if err != nil || date.Day < 0 || date.Day > 31 {
		return Date{}, fmt.Errorf("gtfs: invalid date format, %s", str)
	}

	return date, err
}

// A general time type for the Time gtfs format. The time is measured from "noon
// minus 12h" of the service day (effectively midnight except for days on which
// daylight savings time changes occur. For more information, see the guidelines
// article). For times occurring after midnight, enter the time.Hour can have a
// value greater than 24 for the day on which the trip schedule begins.
//
// The empty value is created with Hour == -1.
//
// Example: 14:30:00 for 2:30PM or 25:35:00 for 1:35AM on the next day.
type Time struct {
	Hour   int
	Minute int
	Second int
}

var TimeEmpty Time = Time{Hour: -1}

func (t Time) IsEmpty() bool {
	return t == TimeEmpty
}

// Parse the time from a string in gtfs format: HH:MM:SS or H:MM:SS
func ParseTime(str string) (Time, error) {
	if len(str) != 8 && len(str) != 7 {
		return Time{}, fmt.Errorf("gtfs: invalid time format, %s", str)
	}

	var err error
	var time Time

	strSplit := strings.Split(str, ":")
	if len(strSplit) != 3 {
		return Time{}, fmt.Errorf("gtfs: invalid time format, %s", str)
	}

	time.Hour, err = strconv.Atoi(strSplit[0])
	if err != nil || time.Hour < 0 {
		return Time{}, fmt.Errorf("gtfs: invalid time format, %s", str)
	}

	time.Minute, err = strconv.Atoi(strSplit[1])
	if err != nil || time.Minute < 0 || time.Minute >= 60 {
		return Time{}, fmt.Errorf("gtfs: invalid time format, %s", str)
	}

	time.Second, err = strconv.Atoi(strSplit[2])
	if err != nil || time.Second < 0 || time.Second >= 60 {
		return Time{}, fmt.Errorf("gtfs: invalid time format, %s", str)
	}

	return time, err
}

//
//
//
//
//

func getString(str string, req bool) (string, error) {
	if str == "" && req {
		return "", errors.New("empty value while required")
	}

	return str, nil
}

func getInt(str string, req bool, empty int) (int, error) {
	if str == "" && req {
		return 0, errors.New("empty value while required")
	}

	if str == "" {
		return empty, nil
	}

	return strconv.Atoi(str)
}

func getFloat(str string, req bool, empty float64) (float64, error) {
	if str == "" && req {
		return 0, errors.New("empty value while required")
	}

	if str == "" {
		return 0, nil
	}

	return strconv.ParseFloat(str, 64)
}

func getDate(str string, req bool) (Date, error) {
	if str == "" && req {
		return DateEmpty, errors.New("empty value while required")
	}

	return ParseDate(str)
}

func getTime(str string, req bool) (Time, error) {
	if str == "" && req {
		return TimeEmpty, errors.New("empty value while required")
	}

	return ParseTime(str)
}
