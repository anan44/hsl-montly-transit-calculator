package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"hsl-transit/transit-calc/hsl"
	"log"
)

type DestinationInput struct {
	Name          string `json:"name"`
	Address       string `json:"address"`
	TimesPerMonth int32  `json:"timesPerMonth"`
}

type TransitInput struct {
	Home         string             `json:"home"`
	Destinations []DestinationInput `json:"destinations"`
}

type TransitOutput struct {
	MonthlyTotalCommute string `json:"monthlyTotalCommute"`
}

func newTransitOutput(commutes *hsl.MonthlyCommutes) TransitOutput {
	return TransitOutput{
		MonthlyTotalCommute: fmt.Sprintf("%v", commutes.TotalDuration()),
	}
}

func HandleLambdaEvent(_ context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var transitInput TransitInput
	bodyBytes := []byte(event.Body)
	err := json.Unmarshal(bodyBytes, &transitInput)
	if err != nil {
		log.Printf("Failded to parse body: %v", event.Body)
		return events.APIGatewayProxyResponse{}, errors.New("failed to unmarshal body")
	}
	home := hsl.Location{Address: transitInput.Home}
	var routes []hsl.Route
	for _, destination := range transitInput.Destinations {
		routes = append(routes, hsl.Route{
			Name:          destination.Name,
			Start:         home,
			End:           hsl.Location{Address: destination.Address},
			TimesPerMonth: destination.TimesPerMonth,
		})
	}
	commutes := hsl.NewMonthlyCommutes(routes)
	transitOutput := newTransitOutput(&commutes)
	output, err := json.Marshal(transitOutput)
	if err != nil {
		log.Printf("Failded to marsal output: %v", transitOutput)
		return events.APIGatewayProxyResponse{}, errors.New("failed to marshal output")
	}
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(output),
	}
	return response, nil
}
