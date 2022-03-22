package dynamodb

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go"
	"log"
	"strings"
)

func CreateLocalClient() *dynamodb.Client{
	cfg,err:=config.LoadDefaultConfig(context.TODO(),
			config.WithRegion("us-east-1"),
			config.WithEndPointResolver(aws.EndpointResolverFunc(
				func(service,region string)(aws.Endpoint,error){
					return aws.Endpoint{URL:"http://localhost:8000"},nil
				}
			)),
			config.WithCredentialsProvider(credentials.StaticCredentials)
				
)
}
