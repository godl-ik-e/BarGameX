package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/godl-ik-e/BarGameX/bar"
	"github.com/godl-ik-e/BarGameX/db"
	"github.com/godl-ik-e/BarGameX/util"
)

type Response struct {
	Bars []*bar.Bar
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("Request: %v", req)
	listOfBars, _ := db.ListBars()
	log.Println(listOfBars)

	// Success HTTP response
	body, err := json.Marshal(&Response{
		Bars: listOfBars,
	})

	if err != nil {
		log.Printf("Error: %", err)
	}

	log.Printf("Body: %v", string(body))
	return util.ResponseToGateway(http.StatusOK, string(body))
}

func main() {
	lambda.Start(Handler)
}
