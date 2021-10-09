package hsl

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestFillCoordinates(t *testing.T) {
	AddressSearchDoFunc = func(req *http.Request) (*http.Response, error) {
		body := []byte(`{"features":[{"geometry":{"type":"Point","coordinates":[24.9,60.1]}}]}`)
		response := http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(body)),
		}
		return &response, nil
	}
	AddressSearchClient = AddressSearchMockClient{}

	location := Location{Address: "Somewhere"}
	location.fillCoordinates()
	got := location.Coordinates
	want := Coordinates{Longitude: 24.9, Latitude: 60.1}
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestEstimateAllRoutesSingleSuccess(t *testing.T) {
	AddressSearchDoFunc = func(req *http.Request) (*http.Response, error) {
		body := []byte(`{"features":[{"geometry":{"type":"Point","coordinates":[24.9,60.1]}}]}`)
		response := http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(body)),
		}
		return &response, nil
	}
	AddressSearchClient = AddressSearchMockClient{}
	TravelDurationDoFunc = func(req *http.Request) (*http.Response, error) {
		body := []byte(`{"data":{"p1":{"itineraries":[{"duration":836}]},"p2":{"itineraries":[{"duration":836}]},"p3":{"itineraries":[{"duration":836}]}}}`)
		response := http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(body)),
		}
		return &response, nil
	}
	TravelDurationClient = TravelDurationMockClient{}

	monthlyCommutes := MonthlyCommutes{Routes: []Route{
		{
			Name:          "TestRoute",
			Start:         Location{Address: "Rautatientori"},
			End:           Location{Address: "One Pint Pub"},
			TimesPerMonth: 1,
		},
	}}
	monthlyCommutes.estimateAllRoutes()
	got := monthlyCommutes.Routes[0].TravelDuration
	want := time.Duration(836) * time.Second
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
