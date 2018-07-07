package main

import (
	"fmt"
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

	putBar(bar)

	returnString, _ := json.Marshal(Response{Message: fmt.Sprintf("Your bar name is %s and has %d female Waitresses!", bar.Name, bar.NumberOfFemaleWaitresses)})

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(returnString),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
