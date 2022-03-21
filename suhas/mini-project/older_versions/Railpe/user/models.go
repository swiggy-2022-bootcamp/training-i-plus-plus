package user

import (
	//"time"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type User struct {
	UUID  string `dynamodbav:"uuid"`
	Name  string `dynamodbav:"name"`
	Email string `dynamodbav:"email"`
}

type UserStorer interface {
	Insert(ctx context.Context, user User) error
	// 	Find(ctx context.Context, uuid string) (User, error)
	// 	Delete(ctx context.Context, uuid string) error
	// 	Update(ctx context.Context, user User) error
}

// type UserStorage struct {
// 	timeout time.Duration
// 	client  *dynamodb.DynamoDB
// }

func AddUser(clnt *dynamodb.Client, u User) error {
	av, err := attributevalue.MarshalMap(u)
	if err != nil {
		fmt.Errorf("failed to marshal Record, %w", err)
		return err
	}

	fmt.Println(av, u)

	out, err := clnt.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("users"),
		Item:      av,
	})

	fmt.Println(out, err)
	return err
}

func ReadUser(clnt *dynamodb.Client, id string) (user User, e error) {
	key := struct {
		UUID string `dynamodbav:"uuid"`
	}{UUID: id}
	u := User{}
	avs, err := attributevalue.MarshalMap(key)
	if err != nil {
		fmt.Errorf("failed to marshal Record, %w", err)
		return u, err
	}
	out, err := clnt.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("users"),
		Key:       avs,
	})
	if err != nil {
		fmt.Errorf("Failed to get item, %w", err)
		return u, err
	}
	fmt.Println(out.Item)
	var readUser User
	if err := attributevalue.UnmarshalMap(out.Item, &readUser); err != nil {
		fmt.Errorf("Failed to unmarshal Record, %w", err)
		return u, err
	}
	return readUser, nil
}

func UpdateUser(clnt *dynamodb.Client, id string, useremail string) error {

	key := struct {
		UUID string `dynamodbav:"uuid"`
	}{UUID: id}

	avs, err := attributevalue.MarshalMap(key)
	if err != nil {
		fmt.Errorf("failed to marshal Record, %w", err)
		return err
	}

	out, err := clnt.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName:        aws.String("users"),
		Key:              avs,
		UpdateExpression: aws.String("SET email = :u"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":u": &types.AttributeValueMemberS{Value: useremail},
		},
	})
	fmt.Println(err)
	fmt.Println(out.Attributes)
	return err
}

func DeleteUser(clnt *dynamodb.Client, id string) error {

	key := struct {
		UUID string `dynamodbav:"uuid"`
	}{UUID: id}

	avs, err := attributevalue.MarshalMap(key)
	if err != nil {
		fmt.Errorf("failed to marshal Record, %w", err)
		return err
	}

	out, err := clnt.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String("users"),
		Key:       avs,
	})
	fmt.Println(out.Attributes)
	return err
}
