package internal

import "time"

type Time struct {
	Hour    int
	Minutes int
}

type Frequency struct {
	Weekdays []time.Weekday
	Times    []Time
}

type RouteStop struct {
	Id                  string
	Name                string
	TimeToArriveSeconds int
}

type Route struct {
	Stops     []RouteStop
	Frequency Frequency
}

type Line struct {
	Id     string
	Name   string
	Routes []Route
}
