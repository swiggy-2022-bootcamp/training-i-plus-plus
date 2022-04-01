package repository

import (
	"context"
	"fmt"
	"srctc/database"
	"srctc/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collectionAuth = new(mongo.Collection)
)

func init() {
	var DB *mongo.Client = database.ConnectDB()
	collectionAuth = DB.Database("golangAPI").Collection("signup")
}

type AuthRepository struct{}

func (ath AuthRepository) Create(newUser models.SignUp) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionAuth.InsertOne(ctx, &newUser)
	if err == nil {
		fmt.Println("Inserted a single document: ", result.InsertedID)
	}

	return result.InsertedID, err
}

func (ath AuthRepository) Read(username string) (models.SignUp, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var regUser models.SignUp
	err := collectionAuth.FindOne(ctx, bson.M{"username": username}).Decode(&regUser)
	return regUser, err
}
