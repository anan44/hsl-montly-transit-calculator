package hsl

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func (m MockClient) Do(req *http.Request) (*http.Response, error) {
	return DoFunc(req)
}

func TestGetTravelDuration(t *testing.T) {
	DoFunc = func(req *http.Request) (*http.Response, error) {
		body := []byte(`{"data":{"p1":{"itineraries":[{"duration":836}]},"p2":{"itineraries":[{"duration":836}]},"p3":{"itineraries":[{"duration":836}]}}}`)
		response := http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(body)),
		}
		return &response, nil
	}
	Client = MockClient{}
	route := Route{
		Name:           "TestRoute",
		Start:          Location{Address: "Leonkatu", Coordinates: Coordinates{Longitude: 24.98, Latitude: 60.18}},
		End:            Location{Address: "Rautatientori", Coordinates: Coordinates{Longitude: 24.93, Latitude: 60.17}},
		TimesPerMonth:  1,
	}
	route.getTravelDuration("2021-10-12")

	got := route.TravelDuration
	want := time.Duration(836) * time.Second
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
