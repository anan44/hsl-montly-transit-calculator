package hsl

import (
	"encoding/json"
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

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

func addressToCoordinates(address string) Coordinates {
	uri := "http://api.digitransit.fi/geocoding/v1/search"
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	q := url.Values{}
	q.Add("text", address)
	q.Add("size", "1")
	req.URL.RawQuery = q.Encode()

	res, err := Client.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var result coordinatesResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}
	if len(result.Features) == 0 {
		panic(err)
	}
	return Coordinates{
		Longitude: result.Features[0].Geometry.Coordinates[0],
		Latitude:  result.Features[0].Geometry.Coordinates[1],
	}
}
