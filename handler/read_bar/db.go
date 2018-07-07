package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/godl-ik-e/BarGameX/bar"
)

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

func getBar(name string) (*bar.Bar, error) {

	input := &dynamodb.GetItemInput{
		TableName: aws.String("aws-go-dep-dev"),
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

	bar := new(bar.Bar)
	err = dynamodbattribute.UnmarshalMap(result.Item, bar)
	if err != nil {
		return nil, err
	}

	return bar, nil
}
