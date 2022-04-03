package repository

import (
	"context"
	"time"
	"userService/config"
	"userService/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collectionTrainName = "trains"
	collectionTrain     = new(mongo.Collection)
)

func init() {
	var DB *mongo.Client = config.ConnectDB()
	collectionTrain = DB.Database("golangAPI").Collection(collectionTrainName)
}

type TrainRepository struct{}

func (trn TrainRepository) ReadAll() ([]models.Train, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	var users []models.Train
	defer cancel()
	results, err := collectionTrain.Find(ctx, bson.M{})
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleTrain models.Train
		if err = results.Decode(&singleTrain); err != nil {
			errLog(err)
			return users, err
		}

		users = append(users, singleTrain)
	}
	return users, nil
}
