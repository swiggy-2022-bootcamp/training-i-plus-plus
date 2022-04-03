package repository

import (
	"authService/config"
	log "authService/logger"
	"authService/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	errLog = log.InfoLogger.Println
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
	if err != nil {
		errLog(err)
	}
	return result.InsertedID, err
}

func (ath AuthRepository) Read(username string) (models.Register, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var regUser models.Register
	err := collectionAuth.FindOne(ctx, bson.M{"username": username}).Decode(&regUser)
	if err != nil {
		errLog(err)
	}
	return regUser, err
}
