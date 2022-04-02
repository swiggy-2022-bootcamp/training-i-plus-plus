package repository

import (
	"TrainService/database"
	"TrainService/logger"
	"TrainService/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collectionTrainName = "trains"
	collectionTrain     = new(mongo.Collection)
	logger6             = logger.NewLoggerService("train_repository")
)

func init() {
	var DB *mongo.Client = database.ConnectDB()
	collectionTrain = DB.Database("golangAPI").Collection(collectionTrainName)
}

type TrainRepository struct{}

func (trn TrainRepository) Create(newTrain models.Train) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionTrain.InsertOne(ctx, &newTrain)
	if err == nil {
		logger6.Log("Created a new train journy: " + newTrain.Source + " to " + newTrain.Destination)
		// fmt.Println("Inserted a single document: ", result.InsertedID)
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
	//update := bson.M{"source": updateTrain.source, "destination": updateTrain.destination}
	updatebson := bson.M{}
	if updateTrain.Source != "" {
		updatebson["source"] = updateTrain.Source
	}
	if updateTrain.Destination != "" {
		updatebson["destination"] = updateTrain.Destination
	}
	result, err := collectionTrain.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": updatebson})
	if err == nil {
		logger6.Log("Updated train journey: " + updateTrain.Source + " to " + updateTrain.Destination)
		// fmt.Println("Updated a single document: ", result.UpsertedID)
	}
	return result.UpsertedID, err
}

func (trn TrainRepository) Delete(objId primitive.ObjectID) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionTrain.DeleteOne(ctx, bson.M{"_id": objId})
	if err == nil {
		logger6.Log("Deleted a train journey: " + objId.Hex())
		// fmt.Println("Updated a single document: ", result.DeletedCount)
	}
	return result.DeletedCount, err
}

func (trn TrainRepository) ReadAll() ([]models.Train, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	var users []models.Train
	defer cancel()

	results, err := collectionTrain.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
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
