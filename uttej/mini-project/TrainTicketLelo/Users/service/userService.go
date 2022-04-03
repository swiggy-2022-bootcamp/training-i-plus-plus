package service

import (
	"Users/config"
	errors "Users/errors"
	"Users/middleware"
	models "Users/model"
	"context"
	"encoding/json"
	"io"
	"log"
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
	client, _ = mongo.NewClient(options.Client().ApplyURI(mongoURL))
	userCollection = client.Database("TrainTicketLelo").Collection("Users")
}

func LogInUser(login models.Login) (jwtToken string, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	result := userCollection.FindOne(ctx, bson.M{"username": login.UserName})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		return "", errors.UnauthorizedError()
	}

	var user models.User
	result.Decode(&user)

	err = middleware.ComparePassword(user.Password, login.Password)
	if err != nil {
		return "", errors.UnauthorizedError()
	}

	jwtToken, _ = middleware.GenerateJWT(user.Id.Hex(), user.Role)
	return
}

func CreateUser(body *io.ReadCloser) (insertedId string, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	var newUser models.User
	json.NewDecoder(*body).Decode(&newUser)

	count, err := userCollection.CountDocuments(ctx, bson.M{"username": newUser.UserName})
	if err != nil {
		return "", err
	}
	if count > 0 {
		return "", errors.UserAlreadyExists()
	}

	hashedPassword, err := middleware.HashMyPassword(newUser.Password)
	if err != nil {
		log.Panic(err)
	}

	newUser.Password = string(hashedPassword)

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := userCollection.InsertOne(ctx, newUser)

	insertedId = result.InsertedID.(primitive.ObjectID).Hex()

	return insertedId, nil
}

func GetAllUsers() (allUsers []models.User) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, _ := userCollection.Find(ctx, bson.M{})

	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		allUsers = append(allUsers, user)
	}
	return
}

func GetUserById(userId string) (userRetrieved *models.User, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

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

func UpdateUserById(userId string, body *io.ReadCloser) (userRetrieved *models.User, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	updatedUser := &models.User{}
	unmarshalErr := json.NewDecoder(*body).Decode(&updatedUser)
	if unmarshalErr != nil {
		return nil, errors.UnmarshallError()
	}

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

	str := "User Deleted"
	successMessage = &str
	return
}
