package repository

import (
	"context"
	"fmt"
	"gin-mongo-api/config"
	"gin-mongo-api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collectionAuthName = "register"
	collectionAuth     = new(mongo.Collection)
)

func init() {
	var DB *mongo.Client = config.ConnectDB()
	collectionAuth = DB.Database("golangAPI").Collection(collectionAuthName)
}

type AuthRepository struct{}

func (ath AuthRepository) Insert(newReg models.Register) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionAuth.InsertOne(ctx, &newReg)
	if err == nil {
		fmt.Println("Inserted a single document: ", result.InsertedID)
	}
	return result.InsertedID, err
}

func (ath AuthRepository) Read(username string) (models.Register, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var regUser models.Register
	err := collectionAuth.FindOne(ctx, bson.M{"username": username}).Decode(&regUser)
	return regUser, err
}
