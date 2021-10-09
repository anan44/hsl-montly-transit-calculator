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
	Data map[string]struct {
		Itineraries []struct {
			Duration int64
		}
	}
}

var (
	TravelDurationClient HTTPClient
)

func singleTravelDurationPlan(planName string, start Coordinates, end Coordinates, date string, hour string) string {
	plan := fmt.Sprintf(`
%s: plan(
	from: {lat: %f, lon: %f}
	to: {lat: %f, lon: %f}
	numItineraries: 1
	date: "%s"
	time: "%s:00"
) {
	itineraries {
	duration
	}
}`, planName, start.Latitude, start.Longitude, end.Latitude, end.Longitude, date, hour)
	return plan
}

func (r *Route) travelDurationApiCall(date string) routeResponse {
	uri := "https://api.digitransit.fi/routing/v1/routers/hsl/index/graphql"
	queryJsonData := map[string]string{
		"query": fmt.Sprintf(`{%s %s %s}`,
			singleTravelDurationPlan("p1", r.Start.Coordinates, r.End.Coordinates, date, "08"),
			singleTravelDurationPlan("p2", r.Start.Coordinates, r.End.Coordinates, date, "17:45"),
			singleTravelDurationPlan("p3", r.Start.Coordinates, r.End.Coordinates, date, "22:20")),
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
	resp, err := TravelDurationClient.Do(req)
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
	return result
}

func (r *Route) getTravelDuration(date string) {
	result := r.travelDurationApiCall(date)
	var totalSeconds int64
	for _, x := range result.Data {
		if len(x.Itineraries) == 0 {
			panic("No Itineraries found")
		}
		totalSeconds += x.Itineraries[0].Duration
	}
	average := totalSeconds / int64(len(result.Data))
	duration := time.Duration(average) * time.Second
	r.TravelDuration = duration
}
