package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

func getBar(name string) (*Bar, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Bars"),
		Key: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String(name),
			},
		},
	}

	result, err := db.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	bar := new(Bar)
	err = dynamodbattribute.UnmarshalMap(result.Item, bar)
	if err != nil {
		return nil, err
	}

	return bar, nil
}
