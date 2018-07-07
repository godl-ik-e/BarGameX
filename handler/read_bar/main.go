package main

import (
	"net/http"

	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Bar struct {
	Name                     string `json:"Barname"`
	NumberOfFemaleWaitresses int    `json:"Number of female Waitresses"`
}

type Response struct {
	Message string `json:"message"`
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("Request: %v", req)

	bar, _ := getBar("Sugar Ray")

	returnString, _ := json.Marshal(bar)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(returnString),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
