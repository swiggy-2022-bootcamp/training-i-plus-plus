package kafka

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"train_service/configs"
	"train_service/models"

	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"

	kafka "github.com/segmentio/kafka-go"
)

const (
	topic         = "test"
	brokerAddress = "localhost:9092"
)
var trainCollection *mongo.Collection = configs.GetCollection(configs.DB, "train_details")

func UpdateSeatsConsumer() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})

	for {

		m, err := r.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}

		strArr := strings.Split(string(m.Value), " ")
		
		if len(strArr) == 3 {
			if strArr[0] == "Book" {
				UpdateSeats(strArr[1], strArr[2], false)
			} else {
				UpdateSeats(strArr[1], strArr[2], true)
			}
		}
		
	}
}

func UpdateSeats(trainId string, seats string, isIncrease bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var train models.Train
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(trainId)

	err := trainCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&train)
	if err != nil {
		//Log Error
			return
	}

	intVar, _ := strconv.ParseInt(seats, 10, 64)

	var availableSeats int64
	if isIncrease {
		availableSeats = train.AvailableSeats + intVar
	} else {
		availableSeats = train.AvailableSeats - intVar
	}

	update := bson.M{
		"availableseats": availableSeats,
	}
	result, err := trainCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		// logg error
		// c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	fmt.Print(result)
	//logg the success result
	// c.JSON(http.StatusOK, responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedTrain}})
}