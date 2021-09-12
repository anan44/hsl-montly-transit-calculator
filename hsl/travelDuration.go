package hsl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type routeResponse struct {
	Data struct {
		Plan struct {
			Itineraries []struct {
				Duration int64 `json:"duration"`
			} `json:"itineraries"`
		} `json:"plan"`
	} `json:"data"`
}

func singleTravelDurationPlan(start Coordinates, end Coordinates, date string, hour string) string {
	plan := fmt.Sprintf(`
			  plan(
				 from: {lat: %f, lon: %f}
				 to: {lat: %f, lon: %f}
				 numItineraries: 1
                 date: "%s"
				 time: "%s:00:00"
			  ) {
				 itineraries {
                   duration
				   legs {
					  startTime
					  endTime
					  route {
						shortName
						longName
					  }
					  mode
					  distance
				   }
				 }
			  }`, start.Latitude, start.Longitude, end.Latitude, end.Longitude, date, hour)
	return plan
}

func (r *Route)getTravelDuration(date string, hour string) {
	uri := "https://api.digitransit.fi/routing/v1/routers/hsl/index/graphql"
	queryJsonData := map[string]string{
		"query": fmt.Sprintf(`
			{
				%s
			}
		`, singleTravelDurationPlan(r.Start.Coordinates, r.End.Coordinates, date, hour)),
	}
	jsonBody, err := json.Marshal(queryJsonData)

	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var result routeResponse
	err = json.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}
	if len(result.Data.Plan.Itineraries) == 0 {
		panic(err)
	}
	seconds := result.Data.Plan.Itineraries[0].Duration
	duration := time.Duration(seconds) * time.Second
	r.TravelDuration = duration
}

