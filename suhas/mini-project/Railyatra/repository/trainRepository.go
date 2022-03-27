package repository

import (
	"context"
	"fmt"
	"gin-mongo-api/config"
	"gin-mongo-api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collectionTrainName = "availticket"
	collectionTrain     = new(mongo.Collection)
)

func init() {
	var DB *mongo.Client = config.ConnectDB()
	collectionTrain = DB.Database("golangAPI").Collection(collectionTrainName)
}

type TrainRepository struct{}

func (trn TrainRepository) Insert(newTrain models.Train) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionTrain.InsertOne(ctx, &newTrain)
	if err == nil {
		fmt.Println("Inserted a single document: ", result.InsertedID)
	}
	return result.InsertedID, err
}

func (trn TrainRepository) Read(objId primitive.ObjectID) (models.Train, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var user models.Train
	err := collectionTrain.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	return user, err
}

func (trn TrainRepository) Update(updateTrain models.Train, objId primitive.ObjectID) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//update := bson.M{"station1": updateTrain.Station1, "station2": updateTrain.Station2}
	updatebson := bson.M{}
	if updateTrain.Station1 != "" {
		updatebson["station1"] = updateTrain.Station1
	}
	if updateTrain.Station2 != "" {
		updatebson["station2"] = updateTrain.Station2
	}
	result, err := collectionTrain.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": updatebson})
	if err == nil {
		fmt.Println("Updated a single document: ", result.UpsertedID)
	}
	return result.UpsertedID, err
}

func (trn TrainRepository) Delete(objId primitive.ObjectID) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionTrain.DeleteOne(ctx, bson.M{"_id": objId})
	if err == nil {
		fmt.Println("Updated a single document: ", result.DeletedCount)
	}
	return result.DeletedCount, err
}

func (trn TrainRepository) ReadAll() ([]models.Train, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	var users []models.Train
	defer cancel()
	results, err := collectionTrain.Find(ctx, bson.M{})
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleTrain models.Train
		if err = results.Decode(&singleTrain); err != nil {
			return users, err
		}

		users = append(users, singleTrain)
	}
	return users, nil
}
