package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

// Add a book record to DynamoDB.
func putBar(bar *Bar) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("aws-go-dep-dev"),
		Item: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String(bar.Name),
			},
			"NumberOfFemaleWaitresses": {
				N: aws.String(fmt.Sprintf("%d", bar.NumberOfFemaleWaitresses)),
			},
		},
	}

	_, err := db.PutItem(input)
	return err
}
