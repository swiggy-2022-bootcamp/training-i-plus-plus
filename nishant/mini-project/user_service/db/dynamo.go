package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/user_service/config"
)

func Connect() (*session.Session, *dynamodb.DynamoDB) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Credentials: credentials.NewStaticCredentials(config.C().DB.ID, config.C().DB.Secret, ""),
			Region:      aws.String(config.C().DB.Region),
			Endpoint:    aws.String(config.C().DB.Endpoint),
		},
	}))

	svc := dynamodb.New(sess)

	return sess, svc
}
