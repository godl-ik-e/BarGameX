package main

import (
	"fmt"
	"net/http"

	"encoding/json"
	"log"

	"github.com/godl-ik-e/BarGameX/bar"
	"github.com/godl-ik-e/BarGameX/util"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("Request: %v", req)

	bar := new(bar.Bar)

	err := json.Unmarshal([]byte(req.Body), bar)

	if err != nil {
		log.Printf("Error Body: %v", req.Body)
		return util.ResponseToGateway(http.StatusBadRequest, req.Body)
	}

	putBar(bar)

	returnString, _ := json.Marshal(Response{Message: fmt.Sprintf("Your bar name is %s and has %d female Waitresses!", bar.Name, bar.NumberOfFemaleWaitresses)})

	return util.ResponseToGateway(http.StatusOK, string(returnString))

}

func main() {
	lambda.Start(Handler)
}
