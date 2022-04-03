package service

import (
	"Trains/config"
	errors "Trains/errors"
	kafka "Trains/kafka"
	models "Trains/models"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var mongoURL string = config.MONGO_URL
var trainCollection *mongo.Collection

func init() {
	client, _ = mongo.NewClient(options.Client().ApplyURI(mongoURL))
	trainCollection = client.Database("TrainTicketLelo").Collection("Trains")
}

func CreateTrain(body *io.ReadCloser) (result *mongo.InsertOneResult) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	var newTrain models.Train
	json.NewDecoder(*body).Decode(&newTrain)

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	result, _ = trainCollection.InsertOne(ctx, newTrain)
	return
}

func GetTrains() (allTrains []models.Train) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := trainCollection.Find(ctx, bson.M{})

	if err != nil {
		fmt.Println(err)
	}

	for cursor.Next(ctx) {
		var train models.Train
		cursor.Decode(&train)
		allTrains = append(allTrains, train)
	}
	return
}

func GetTrainById(trainId string) (trainRetrieved *models.Train, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	objectId, err := primitive.ObjectIDFromHex(trainId)

	if err != nil {
		return nil, errors.MalformedIdError()
	}

	result := trainCollection.FindOne(ctx, bson.M{"_id": objectId})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		return nil, errors.IdNotFoundError()
	}

	result.Decode(&trainRetrieved)
	return
}

func UpdateTrainById(trainId string, body *io.ReadCloser) (trainRetrieved *models.Train, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	var updatedTrain models.Train
	unmarshalErr := json.NewDecoder(*body).Decode(&updatedTrain)
	if unmarshalErr != nil {
		return nil, errors.UnmarshallError()
	}

	return UpdateTrainByIdWorker(trainId, updatedTrain)
}

func UpdateTrainByIdWorker(trainId string, updatedTrain models.Train) (trainRetrieved *models.Train, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	objectId, err := primitive.ObjectIDFromHex(trainId)
	if err != nil {
		return nil, errors.MalformedIdError()
	}

	result, error := trainCollection.UpdateByID(ctx, objectId, bson.M{"$set": updatedTrain})
	if error != nil {
		return nil, errors.InternalServerError()
	}

	if result.MatchedCount == 0 {
		return nil, errors.IdNotFoundError()
	}
	return GetTrainById(trainId)
}

func DeleteTrainbyId(trainId string) (successMessage *string, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	objectId, err := primitive.ObjectIDFromHex(trainId)
	if err != nil {
		return nil, errors.MalformedIdError()
	}

	result, error := trainCollection.DeleteOne(ctx, bson.M{"_id": objectId})

	if error != nil {
		return nil, errors.InternalServerError()
	}

	if result.DeletedCount == 0 {
		return nil, errors.IdNotFoundError()
	}
	msg := "Train Deleted"
	successMessage = &msg
	return
}

func UpdateTicketCount(trainId string, updateCount int) (quantityAfterUpdation *int, err error) {
	trainRetrieved, error := GetTrainById(trainId)

	if error != nil {
		trainError, ok := error.(*errors.TrainError)
		if ok {
			return nil, trainError
		} else {
			fmt.Println("Couldn't Update the Ticket Count")
			return
		}
	}

	trainRetrieved.AvailableTickets += updateCount
	if trainRetrieved.AvailableTickets < 0 {
		ctx, _ := context.WithTimeout(context.Background(), time.Minute*10)
		kafka.Produce(ctx, nil, []byte("trainId: "+trainId+" --- status: Tickets Sold Out (critical)"))

		return nil, errors.OutOfStockError()
	}

	_, err = UpdateTrainByIdWorker(trainId, *trainRetrieved)
	if err != nil {
		return nil, err
	}

	return &trainRetrieved.AvailableTickets, nil
}
