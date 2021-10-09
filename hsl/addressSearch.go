package hsl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type coordinatesResponse struct {
	Features []struct {
		Geometry struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}

var AddressSearchClient HTTPClient

func addressToCoordinates(address string) (Coordinates, error) {
	uri := "http://api.digitransit.fi/geocoding/v1/search"
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	q := url.Values{}
	q.Add("text", address)
	q.Add("size", "1")
	req.URL.RawQuery = q.Encode()

	res, err := AddressSearchClient.Do(req)
	if err != nil {
		return Coordinates{}, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Coordinates{}, err
	}
	var result coordinatesResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return Coordinates{}, err
	}
	if len(result.Features) == 0 {
		return Coordinates{}, fmt.Errorf("no features found from result")
	}
	return Coordinates{
		Longitude: result.Features[0].Geometry.Coordinates[0],
		Latitude:  result.Features[0].Geometry.Coordinates[1],
	}, nil
}
