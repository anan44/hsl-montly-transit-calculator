package hsl

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockClient struct{}

var DoFunc func(req *http.Request) (*http.Response, error)

func (m MockClient) Do(req *http.Request) (*http.Response, error) {
	return DoFunc(req)
}

func TestAddressToCoordinates(t *testing.T) {
	DoFunc = func(req *http.Request) (*http.Response, error) {
		body := []byte(`{"features":[{"geometry":{"type":"Point","coordinates":[24.9,60.1]}}]}`)
		response := http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(body)),
		}
		return &response, nil
	}
	address := "some address"
	Client = MockClient{}
	got := addressToCoordinates(address)
	want := Coordinates{Longitude: 24.9, Latitude: 60.1}
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
