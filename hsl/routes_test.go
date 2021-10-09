package hsl

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestFillCoordinates(t *testing.T) {
	DoFunc = func(req *http.Request) (*http.Response, error) {
		body := []byte(`{"features":[{"geometry":{"type":"Point","coordinates":[24.9,60.1]}}]}`)
		response := http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(body)),
		}
		return &response, nil
	}
	AddressSearchClient = MockClient{}

	location := Location{Address: "Somewhere"}
	location.fillCoordinates()
	got := location.Coordinates
	want := Coordinates{Longitude: 24.9, Latitude: 60.1}
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}