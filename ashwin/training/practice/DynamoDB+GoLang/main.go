package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("./.aws/credentials", "test-account"),
	})
	svc := dynamodb.New(sess)

	req := &dynamodb.DescribeTableInput{
		TableName: aws.String("UserDatabase"),
	}
	result, err := svc.DescribeTable(req)
	if err != nil {
		fmt.Printf("%s", err)
	}
	table := result.Table
	fmt.Printf("done", table)
}
