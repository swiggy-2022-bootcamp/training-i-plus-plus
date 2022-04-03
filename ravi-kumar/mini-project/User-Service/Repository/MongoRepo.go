package repository

import (
	"User-Service/config"
	"User-Service/errors"
	mockdata "User-Service/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var client *mongo.Client
var mongoURL string = config.MONGO_URL
var userCollection *mongo.Collection

func init() {
	// Initialize a new mongo client with options
	client, _ = mongo.NewClient(options.Client().ApplyURI(mongoURL))
	userCollection = client.Database("swiggy_mini").Collection("users")
}

type IMongoDAO interface {
	MongoUserLogin(logInDTO mockdata.LogInDTO) (mockdata.User, error)
	MongoCreateUser(newUser mockdata.User) (insertedId string)
	MongoGetAllUsers() []mockdata.User
	MongoGetUserById(userId primitive.ObjectID) (userRetrieved *mockdata.User, err error)
	MongoUpdateUserById(userId primitive.ObjectID, updatedUser mockdata.User) (userRetrieved *mockdata.User, err error)
	MongoDeleteUserById(userId primitive.ObjectID) (successMessage *string, err error)
}

type MongoDAO struct {
}

func (dao *MongoDAO) MongoUserLogin(logInDTO mockdata.LogInDTO) (mockdata.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	result := userCollection.FindOne(ctx, bson.M{"username": logInDTO.UserName})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		return mockdata.User{}, errors.UnauthorizedError()
	}

	var user mockdata.User
	result.Decode(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logInDTO.Password))
	if err == nil {
		//password has matched
		return user, nil
	}
	return mockdata.User{}, errors.UnauthorizedError()
}

func (dao *MongoDAO) MongoCreateUser(newUser mockdata.User) (insertedId string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	result := userCollection.FindOne(ctx, bson.M{"username": newUser.UserName})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		result, _ := userCollection.InsertOne(ctx, newUser)
		insertedId = result.InsertedID.(primitive.ObjectID).Hex()
		return
	}

	return ""
}

func (dao *MongoDAO) MongoGetAllUsers() (allUsers []mockdata.User) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	cursor, _ := userCollection.Find(ctx, bson.M{})

	for cursor.Next(ctx) {
		var user mockdata.User
		cursor.Decode(&user)
		allUsers = append(allUsers, user)
	}
	return
}

func (dao *MongoDAO) MongoGetUserById(userId primitive.ObjectID) (userRetrieved *mockdata.User, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	result := userCollection.FindOne(ctx, bson.M{"_id": userId})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		return nil, errors.IdNotFoundError()
	}

	result.Decode(&userRetrieved)
	return
}

func (dao *MongoDAO) MongoUpdateUserById(userId primitive.ObjectID, updatedUser mockdata.User) (userRetrieved *mockdata.User, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	result, error := userCollection.UpdateByID(ctx, userId, bson.M{"$set": updatedUser})
	if error != nil {
		return nil, errors.InternalServerError()
	}

	if result.MatchedCount == 0 {
		return nil, errors.IdNotFoundError()
	}

	return dao.MongoGetUserById(userId)
}

func (dao *MongoDAO) MongoDeleteUserById(userId primitive.ObjectID) (successMessage *string, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	result, error := userCollection.DeleteOne(ctx, bson.M{"_id": userId})

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
