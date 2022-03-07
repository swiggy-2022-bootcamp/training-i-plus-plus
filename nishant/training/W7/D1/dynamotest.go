package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {

	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000")},
	)
	if err != nil {
		// Handle Session creation error
	}
	// Create DynamoDB client
	svc := dynamodb.New(sess)

	fmt.Println("created dynamo session ", svc)

}
