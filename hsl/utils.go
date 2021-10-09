package hsl

import (
	"net/http"
	"time"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type TravelDurationMockClient struct{}

var TravelDurationDoFunc func(req *http.Request) (*http.Response, error)

func (m TravelDurationMockClient) Do(req *http.Request) (*http.Response, error) {
	return TravelDurationDoFunc(req)
}

type AddressSearchMockClient struct{}

var AddressSearchDoFunc func(req *http.Request) (*http.Response, error)

func (m AddressSearchMockClient) Do(req *http.Request) (*http.Response, error) {
	return AddressSearchDoFunc(req)
}

func init() {
	TravelDurationClient = &http.Client{}
	AddressSearchClient = &http.Client{}
}

func NextMonday() string {
	today := time.Now()
	var nextMonday time.Time
	currentWeekday := today.Weekday()
	switch currentWeekday {
	case time.Monday:
		nextMonday = today.Add(24 * time.Hour * 7)
	case time.Tuesday:
		nextMonday = today.Add(24 * time.Hour * 6)
	case time.Wednesday:
		nextMonday = today.Add(24 * time.Hour * 5)
	case time.Thursday:
		nextMonday = today.Add(24 * time.Hour * 4)
	case time.Friday:
		nextMonday = today.Add(24 * time.Hour * 3)
	case time.Saturday:
		nextMonday = today.Add(24 * time.Hour * 2)
	case time.Sunday:
		nextMonday = today.Add(24 * time.Hour * 1)
	}
	return nextMonday.Format("2006-01-02")
}
