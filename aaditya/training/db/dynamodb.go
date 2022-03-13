package db

import (
	"context"
	"fmt"
	"log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func CreateLocalClient() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8042"}, nil
			})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "dummy", SecretAccessKey: "dummy", SessionToken: "dummy",
				Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}),
	)
	if err != nil {
		panic(err)
	}

	return dynamodb.NewFromConfig(cfg)
}

func CreateTable(cfg *dynamodb.Client){
	out, err := cfg.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
        AttributeDefinitions: []types.AttributeDefinition{
            {
                AttributeName: aws.String("id"),
                AttributeType: types.ScalarAttributeTypeS,
            },
        },
        KeySchema: []types.KeySchemaElement{
            {
                AttributeName: aws.String("id"),
                KeyType:       types.KeyTypeHash,
            },
        },
        TableName:   aws.String("my-table"),
        BillingMode: types.BillingModePayPerRequest,
    })
    if err != nil {
        panic(err)
    }

    fmt.Println(out)
}

func TableExists(d *dynamodb.Client, name string) bool {
	tables, err := d.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		log.Fatal("ListTables failed", err)
	}
	for _, n := range tables.TableNames {
		fmt.Println(n)
		if n == name {
			return true
		}
	}
	return false
}

func InsertItem(d *dynamodb.Client, item string)  {
	_, err := d.PutItem(context.TODO(), &dynamodb.PutItemInput{
        TableName: aws.String("my-table"),
        Item: map[string]types.AttributeValue{
            "id":    &types.AttributeValueMemberS{Value: item},
        },
    })

    if err != nil {
        panic(err)
    }

    fmt.Printf("Item %s inserted successfully.", item)

}

func GetItem(d *dynamodb.Client, item string){
	out, err := d.GetItem(context.TODO(), &dynamodb.GetItemInput{
        TableName: aws.String("my-table"),
        Key: map[string]types.AttributeValue{
            "id": &types.AttributeValueMemberS{Value: item},
        },
    })

    if err != nil {
        panic(err)
    }

    fmt.Println(out.Item["id"])
}

func DeleteItem(d *dynamodb.Client, item string){
	_, err := d.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
        TableName: aws.String("my-table"),
        Key: map[string]types.AttributeValue{
            "id": &types.AttributeValueMemberS{Value: item},
        },
    })
    if err != nil {
        panic(err)
    }

    fmt.Printf("Item %s deleted successfully.", item)
}

func GetAllItems(d *dynamodb.Client){
	out, err := d.Scan(context.TODO(), &dynamodb.ScanInput{
        TableName: aws.String("my-table"),
    })
    if err != nil {
        panic(err)
    }

    fmt.Println(out.Items)
	for _,element := range out.Items {
		fmt.Println(element["id"])
	}
}