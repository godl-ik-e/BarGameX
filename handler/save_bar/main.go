package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Bar struct {
	Name                     string `json:"Barname"`
	NumberOfFemaleWaitresses int    `json:"Number of female Waitresses"`
}

type Response struct {
	Message string `json:"message"`
}

func Handler(bar Bar) (Response, error) {
	return Response{Message: fmt.Sprintf("Your bar name is %s and has %d female Waitresses!", bar.Name, bar.NumberOfFemaleWaitresses)}, nil
}

func main() {
	lambda.Start(Handler)
}
