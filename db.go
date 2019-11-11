package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"strconv"
)

// Declare a new DynamoDB instance. Note that this is safe for concurrent
// use.
var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-south-1"))

func getItem(id float64) (*employee, error) {
	// Prepare the input for the query.
	numstr := strconv.Itoa(int(id))
    input := &dynamodb.GetItemInput{
        TableName: aws.String("Employee"),
        Key: map[string]*dynamodb.AttributeValue{
            "ID": {
                N: aws.String(numstr),
            },
        },
    }

    // Retrieve the item from DynamoDB. If no matching item is found
    // return nil.
    result, err := db.GetItem(input)
    if err != nil {
        return nil, err
    }
    if result.Item == nil {
        return nil, nil
    }

    // The result.Item object returned has the underlying type
    // map[string]*AttributeValue. We can use the UnmarshalMap helper
    // to parse this straight into the fields of a struct. Note:
    // UnmarshalListOfMaps also exists if you are working with multiple
    // items.
    record := new(employee)
    err = dynamodbattribute.UnmarshalMap(result.Item, record)
    if err != nil {
        return nil, err
    }

    return record, nil
}