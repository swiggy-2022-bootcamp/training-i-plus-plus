package controller

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"usecase/user_crud_dynamo/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Dynamo *dynamodb.DynamoDB
}

func (cont Controller) CreateUser(c *gin.Context) {
	newUser := model.User{}

	if err := c.BindJSON(&newUser); err != nil {
		c.Error(err)
		return
	}
	log.Printf("user %+v", newUser)

	newUser.GeneraterId()

	log.Printf("new user %+v", newUser)

	av, err := dynamodbattribute.MarshalMap(newUser)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
		c.Error(err)
		return
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(model.UserTableName),
	}

	_, err = cont.Dynamo.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
		c.Error(err)
		return
	}
}

func (cont Controller) ReadUser(c *gin.Context) {
	id := c.Param("_id")
	if id == "" {
		c.Error(fmt.Errorf("id not found"))
	}

	result, err := cont.Dynamo.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(model.UserTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"userId": {
				S: aws.String(id),
			},
		},
		ProjectionExpression: model.GetDefaultUserProjection(),
	})

	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	log.Printf("fetch result %+v", result)

	if result.Item == nil {
		c.JSON(404, gin.H{
			"msg": "user not found",
		})
	}

	usr := model.User{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &usr)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	c.JSON(200, usr)
}

func (cont Controller) UpdateUser(c *gin.Context) {
	id := c.Param("_id")
	if id == "" {
		c.Error(fmt.Errorf("id not found"))
	}

	usr := model.User{}
	if err := c.BindJSON(&usr); err != nil {
		c.Error(err)
		return
	}
	log.Printf("user %+v", usr)

	updateExp := getUserUpdateExpression(usr)
	expr, err := expression.NewBuilder().WithUpdate(updateExp).Build()

	if err != nil {
		log.Println("error while creating expression")
		c.Error(err)
		return
	}

	cont.Dynamo.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String(model.UserTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"userId": {
				S: aws.String(id),
			},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	})
}

func (cont Controller) DeleteUser(c *gin.Context) {
	id := c.Param("_id")
	if id == "" {
		c.Error(fmt.Errorf("id not found"))
	}

	res, err := cont.Dynamo.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(model.UserTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"userId": {
				S: aws.String(id),
			},
		},
	})

	log.Printf("delete res %+v", res)

	if err != nil {
		log.Println("delete err ", err)
		c.Error(err)
		return
	}

	c.Status(200)

}

func (cont Controller) ListUser(c *gin.Context) {

	res, err := cont.Dynamo.Scan(&dynamodb.ScanInput{
		TableName: aws.String(model.UserTableName),
	})

	if err != nil {
		log.Println("List err ", err)
		c.Error(err)
		return
	}

	if res.Items == nil {
		c.JSON(404, gin.H{
			"msg": "users not found",
		})
	}

	var result []model.User

	err = dynamodbattribute.UnmarshalListOfMaps(res.Items, &result)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	c.JSON(200, result)
}

func getUserUpdateExpression(user model.User) expression.UpdateBuilder {
	// iterate through all fields and set all fields which are not null
	v := reflect.ValueOf(user)
	t := reflect.TypeOf(&user).Elem()
	var updateExp expression.UpdateBuilder

	for i := 0; i < v.NumField(); i++ {
		vf := v.Field(i)
		if !vf.IsZero() {
			nameTag := t.Field(i).Tag.Get("json")
			nameTag = strings.Split(nameTag, ",")[0]
			updateExp = updateExp.Set(expression.Name(nameTag), expression.Value(vf.String()))
			fmt.Printf("Field: %s\tValue: %v\n", nameTag, vf.String())
		}
	}

	log.Printf("update expression %+v", updateExp)
	return updateExp
}
