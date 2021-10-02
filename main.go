package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"hsl-transit/transit-calc/api"
)

func main() {
	lambda.Start(api.HandleLambdaEvent)
}
