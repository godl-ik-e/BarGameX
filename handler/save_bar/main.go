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

	bar := new(Bar)

	err := json.Unmarshal([]byte(req.Body), bar)

	if err != nil {
		log.Printf("Error Body: %v", req.Body)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       req.Body,
		}, nil
	}

	barString, _ := json.Marshal(bar)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(barString),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
