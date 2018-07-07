package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/godl-ik-e/BarGameX/bar"
)

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

func listBars() ([]bar.Bar, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("aws-go-dep-dev"),
	}
	result, _ := db.Scan(input)

	// Construct todos from response
	var bars []bar.Bar
	for _, i := range result.Items {
		bar := bar.Bar{}
		if err := dynamodbattribute.UnmarshalMap(i, &bar); err != nil {
			fmt.Println("Failed to unmarshal")
			fmt.Println(err)
		}
		bars = append(bars, bar)
	}

	return bars, nil
}
