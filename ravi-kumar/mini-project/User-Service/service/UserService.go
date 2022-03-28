package service

import (
	"User-Service/config"
	errors "User-Service/errors"
	mockdata "User-Service/model"
	"context"
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var mongoURL string = config.MONGO_URL
var userCollection *mongo.Collection

func init() {
	// Initialize a new mongo client with options
	client, _ = mongo.NewClient(options.Client().ApplyURI(mongoURL))
	userCollection = client.Database("swiggy_mini").Collection("users")
}

func CreateUser(body *io.ReadCloser) *mongo.InsertOneResult {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	var newUser mockdata.User
	json.NewDecoder(*body).Decode(&newUser)

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := userCollection.InsertOne(ctx, newUser)

	return result
}

func GetAllUsers() (allUsers []mockdata.User) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, _ := userCollection.Find(ctx, bson.M{})

	for cursor.Next(ctx) {
		var user mockdata.User
		cursor.Decode(&user)
		allUsers = append(allUsers, user)
	}
	return
}

func GetUserById(userId string) (userRetrieved *mockdata.User, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return nil, errors.MalformedIdError()
	}

	result := userCollection.FindOne(ctx, bson.M{"_id": objectId})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		return nil, errors.IdNotFoundError()
	}

	result.Decode(&userRetrieved)
	return
}

func UpdateUserById(userId string, body *io.ReadCloser) (userRetrieved *mockdata.User, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	updatedUser := &mockdata.User{}
	unmarshalErr := json.NewDecoder(*body).Decode(&updatedUser)
	if unmarshalErr != nil {
		return nil, errors.UnmarshallError()
	}

	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.MalformedIdError()
	}

	result, error := userCollection.UpdateByID(ctx, objectId, bson.M{"$set": updatedUser})
	if error != nil {
		return nil, errors.InternalServerError()
	}

	if result.MatchedCount == 0 {
		return nil, errors.IdNotFoundError()
	}

	return GetUserById(userId)
}

func DeleteUserbyId(userId string) (successMessage *string, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.MalformedIdError()
	}

	result, error := userCollection.DeleteOne(ctx, bson.M{"_id": objectId})

	if error != nil {
		return nil, errors.InternalServerError()
	}

	if result.DeletedCount == 0 {
		return nil, errors.IdNotFoundError()

	}

	str := "user deleted"
	successMessage = &str
	return
}
