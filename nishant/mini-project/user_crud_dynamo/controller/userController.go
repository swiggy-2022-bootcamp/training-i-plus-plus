package controller

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/user_crud_dynamo/model"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/user_crud_dynamo/producer"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/user_crud_dynamo/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Dynamo   *dynamodb.DynamoDB
	Producer *producer.Producer
}

// user Structs for req - response

// User request info
// @Description User information
type UserUpdateRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"username,omitempty"`
}

// User Response
// @Description User information
type UserResponse struct {
	UserId string `json:"userId,omitempty"`
	Email  string `json:"email,omitempty"`
	Name   string `json:"username,omitempty"`
}

// CreateUser godoc
// @Summary User Sign-Up
// @Description register new user
// @Tags Users
// @Param   user      body UserUpdateRequest true  "user info"
// @Accept  json
// @Success 200
// @Failure 500
// @Router /user [post]
func (cont Controller) CreateUser(c *gin.Context) {
	usr := UserUpdateRequest{}

	if err := c.BindJSON(&usr); err != nil {
		c.Error(err)
		return
	}
	log.Printf("user %+v", usr)

	if !utils.IsEmailValid(usr.Email) {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "invalid email",
		})
		return
	}

	//newUser.GeneraterId()
	newUser := model.NewUser(usr.Name, usr.Email, utils.HashPass(usr.Password))

	log.Printf("new user %+v", newUser)

	av, err := dynamodbattribute.MarshalMap(newUser)
	if err != nil {
		log.Fatalf("Got error marshalling new user item: %s", err)
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

	//send notification to user
	cont.Producer.SendWelcomeEmail(*newUser)

	c.Set("userId", newUser.UserId)
	c.Next()
}

// ReadUser godoc
// @Summary Get User by id
// @Tags Users
// @Param _id path string true "UserId"
// @Success 200 {object} UserResponse
// @Failure 500
// @Failure 400
// @Failure 401
// @Router /user/{_id} [get]
// @Security ApiKeyAuth
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

	usr := UserResponse{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &usr)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	c.JSON(200, usr)
}

// UpdateUser godoc
// @Summary Update User by id
// @Tags Users
// @Param _id path string true "UserId"
// @Param   user      body UserUpdateRequest true  "user info"
// @Success 200
// @Failure 500
// @Failure 400
// @Failure 401
// @Router /user/{_id} [patch]
// @Security ApiKeyAuth
func (cont Controller) UpdateUser(c *gin.Context) {
	id := c.Param("_id")
	if id == "" {
		c.Error(fmt.Errorf("id not found"))
	}

	if id != c.GetString("userId") {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "Can only Update logged in user",
		})
		return
	}

	usr := UserUpdateRequest{}
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

// DeleteUser godoc
// @Summary Delete User by id
// @Tags Users
// @Param _id path string true "UserId"
// @Success 200
// @Failure 500
// @Failure 400
// @Failure 401
// @Router /user/{_id} [delete]
// @Security ApiKeyAuth
func (cont Controller) DeleteUser(c *gin.Context) {
	id := c.Param("_id")
	if id == "" {
		c.Error(fmt.Errorf("id not found"))
	}

	if id != c.GetString("userId") {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "Can only delete logged in user",
		})
		return
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

// ListUser godoc
// @Summary List all users
// @Tags Users
// @Success 200 {array} UserResponse
// @Failure 500
// @Failure 400
// @Failure 401
// @Router /user [get]
// @Security ApiKeyAuth
func (cont Controller) ListUser(c *gin.Context) {

	res, err := cont.Dynamo.Scan(&dynamodb.ScanInput{
		TableName:            aws.String(model.UserTableName),
		ProjectionExpression: model.GetDefaultUserProjection(),
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

func getUserUpdateExpression(user UserUpdateRequest) expression.UpdateBuilder {
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
