package main

import (
	"net/http"

	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/godl-ik-e/BarGameX/util"
)

type Response struct {
	Message string
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("Request: %v", req)

	bar, err := getBar("Sugar Ray")

	if err != nil {
		log.Println(err)
		return util.ResponseToGateway(http.StatusNotFound, "Keine Daten gefunden")
	}

	returnString, _ := json.Marshal(bar)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(returnString),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
