package hsl

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)


func TestGetTravelDuration(t *testing.T) {
	TravelDurationDoFunc = func(req *http.Request) (*http.Response, error) {
		body := []byte(`{"data":{"p1":{"itineraries":[{"duration":836}]},"p2":{"itineraries":[{"duration":836}]},"p3":{"itineraries":[{"duration":836}]}}}`)
		response := http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(body)),
		}
		return &response, nil
	}
	TravelDurationClient = TravelDurationMockClient{}
	route := Route{
		Name:          "TestRoute",
		Start:         Location{Address: "Leonkatu", Coordinates: Coordinates{Longitude: 24.98, Latitude: 60.18}},
		End:           Location{Address: "Rautatientori", Coordinates: Coordinates{Longitude: 24.93, Latitude: 60.17}},
		TimesPerMonth: 1,
	}
	route.getTravelDuration("2021-10-12")

	got := route.TravelDuration
	want := time.Duration(836) * time.Second
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestSingleTravelDurationPlan(t *testing.T) {

	start := Coordinates{Longitude: 24.98, Latitude: 60.18}
	end := Coordinates{Longitude: 24.93, Latitude: 60.17}
	got := singleTravelDurationPlan("TestPlan", start, end, "2021-10-12", "12:15")
	want := `
TestPlan: plan(
	from: {lat: 60.180000, lon: 24.980000}
	to: {lat: 60.170000, lon: 24.930000}
	numItineraries: 1
	date: "2021-10-12"
	time: "12:15:00"
) {
	itineraries {
	duration
	}
}`
	if got != want {
		t.Errorf("want:\n%v\ngot:\n%v", want, got)
	}
}
