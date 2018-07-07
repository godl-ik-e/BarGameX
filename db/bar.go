package db

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/godl-ik-e/BarGameX/bar"
)

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

// Add a book record to DynamoDB.
func PutBar(bar *bar.Bar) error {
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

func GetBar(name string) (*bar.Bar, error) {

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

func ListBars() ([]bar.Bar, error) {
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
