package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/godl-ik-e/BarGameX/bar"
	"github.com/godl-ik-e/BarGameX/util"
)

type Response struct {
	bars []bar.Bar
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("Request: %v", req)
	var bars []bar.Bar
	bars, _ = listBars()

	// Success HTTP response
	body, _ := json.Marshal(&Response{
		bars: bars,
	})
	return util.ResponseToGateway(http.StatusOK, string(body))
}

func main() {
	lambda.Start(Handler)
}
