package hsl

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAddressToCoordinatesSuccess(t *testing.T) {
	AddressSearchDoFunc = func(req *http.Request) (*http.Response, error) {
		body := []byte(`{"features":[{"geometry":{"type":"Point","coordinates":[24.9,60.1]}}]}`)
		response := http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(body)),
		}
		return &response, nil
	}
	address := "some address"
	AddressSearchClient = AddressSearchMockClient{}
	got, err := addressToCoordinates(address)
	if err != nil {
		t.Errorf("addressToCoordinates returned an error: %v", err)
	}
	want := Coordinates{Longitude: 24.9, Latitude: 60.1}
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestAddressToCoordinatesErrorNoFeatures(t *testing.T) {
	AddressSearchDoFunc = func(req *http.Request) (*http.Response, error) {
		body := []byte(`{"features":[]}`)
		response := http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader(body)),
		}
		return &response, nil
	}
	address := "some address"
	AddressSearchClient = AddressSearchMockClient{}
	_, err := addressToCoordinates(address)
	if err == nil {
		t.Error("addressToCoordinates did not return error when no features was returned")
	}
}
