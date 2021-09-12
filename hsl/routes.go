package hsl

import (
	"sync"
	"time"
)

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

type Location struct {
	Address string
	Coordinates Coordinates
}

type Route struct {
	Name string
	Start Location
	End Location
	TimesPerMonth int32
	TravelDuration time.Duration
}

type MonthlyCommutes struct {
	Routes []Route
}

func NewRoute(name string, start Location, end Location, timesPerMonth int32) Route {
	return Route{Name: name, Start: start, End: end, TimesPerMonth: timesPerMonth}
}

func (r *Route) Estimate(wg *sync.WaitGroup) {
	r.fillLocations()
	r.getTravelDuration("2021-09-15", "12")
	wg.Done()
}

func (r *Route) fillLocations() {
	r.Start.fillCoordinates()
	r.End.fillCoordinates()
}

func (r *Location) fillCoordinates() {
	coordinates := addressToCoordinates(r.Address)
	r.Coordinates = coordinates
}

func (mc *MonthlyCommutes) TotalDuration() time.Duration{
	var total float64
	for _, r := range mc.Routes {
		total += r.TravelDuration.Seconds() * float64(r.TimesPerMonth) * 2
	}
	return time.Duration(total) * time.Second
}
