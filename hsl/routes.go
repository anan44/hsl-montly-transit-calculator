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

func NewRoute(name string, start Location, end Location) Route {
	return Route{Name: name, Start: start, End: end}
}

func (r *Route) Calculate(wg *sync.WaitGroup) {
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