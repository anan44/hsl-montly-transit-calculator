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
	Address     string
	Coordinates Coordinates
}

type Route struct {
	Name           string
	Start          Location
	End            Location
	TimesPerMonth  int32
	TravelDuration time.Duration
}

type MonthlyCommutes struct {
	Routes []Route
}

func NewMonthlyCommutes(routes []Route) MonthlyCommutes {
	mc := MonthlyCommutes{routes}
	mc.estimateAllRoutes()
	return mc
}

func (r *Route) estimate(wg *sync.WaitGroup) {
	r.fillLocations()
	r.getAverageTravelDuration(NextMonday())
	wg.Done()
}

func (r *Route) fillLocations() {
	r.Start.fillCoordinates()
	r.End.fillCoordinates()
}

func (r *Location) fillCoordinates() {
	coordinates, err := addressToCoordinates(r.Address)
	if err != nil {
		panic(err)
	}
	r.Coordinates = coordinates
}

func (mc *MonthlyCommutes) TotalDuration() time.Duration {
	var total float64
	for _, r := range mc.Routes {
		total += r.TravelDuration.Seconds() * float64(r.TimesPerMonth) * 2
	}
	return time.Duration(total) * time.Second
}

type RouteMonthlyDuration struct {
	Name            string
	DurationPerTrip time.Duration
	TotalDuration   time.Duration
}

func (mc *MonthlyCommutes) TotalDurationByRoute() []RouteMonthlyDuration {
	var allRoutes []RouteMonthlyDuration
	for _, r := range mc.Routes {
		monthlyTravel := r.TravelDuration.Seconds() * float64(r.TimesPerMonth) * 2
		allRoutes = append(allRoutes, RouteMonthlyDuration{
			Name:     r.Name,
			TotalDuration: time.Duration(monthlyTravel) * time.Second,
			DurationPerTrip: r.TravelDuration,
		})
	}
	return allRoutes
}

func (mc *MonthlyCommutes) estimateAllRoutes() {
	wg := sync.WaitGroup{}
	for i := range mc.Routes {
		wg.Add(1)
		go mc.Routes[i].estimate(&wg)
	}
	wg.Wait()
}
