package util

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func ResponseToGateway(StatusCode int, body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       body,
	}, nil
}
