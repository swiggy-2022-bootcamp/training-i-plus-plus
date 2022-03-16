package controller

import (
	"log"

	"usecase/user_crud_dynamo/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Dynamo *dynamodb.DynamoDB
}

func (cont Controller) CreateUser(c *gin.Context) {
	newUser := model.User{}
	if err := c.Bind(&newUser); err != nil {
		c.Error(err)
		return
	}

	newUser.GeneraterId()

	log.Println("new user ", newUser)

	av, err := dynamodbattribute.MarshalMap(newUser)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(model.UserTableName),
	}

	_, err = cont.Dynamo.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

}
