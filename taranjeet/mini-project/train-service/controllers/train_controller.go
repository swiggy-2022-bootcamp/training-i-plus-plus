package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/taran1515/crud/configs"
	"github.com/taran1515/crud/models"
	"github.com/taran1515/crud/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"time"
)

var trainCollection = configs.GetCollection(configs.DB, "train")
var validate *validator.Validate = validator.New()

func CreateTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
		var train models.Train
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&train); err != nil {
			c.JSON(http.StatusBadRequest, responses.TrainResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&train); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.TrainResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		//var availableSeat int[] = []

		newTrain := models.Train{
			TrainId:       primitive.NewObjectID(),
			Name:          train.Name,
			Source:        train.Source,
			Destination:   train.Destination,
			ArrivalTime:   train.ArrivalTime,
			DepartureTime: train.DepartureTime,
			NumberOfSeats: train.NumberOfSeats,
			ReservedSeats: train.ReservedSeats,
			TrainNumber:   train.TrainNumber,
		}

		fmt.Println("train", newTrain)

		result, err := trainCollection.InsertOne(ctx, newTrain)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.TrainResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetATrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		trainId := c.Param("trainId")
		var train models.Train
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(trainId)

		err := trainCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&train)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": train}})
	}
}

func EditATrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		trainId := c.Param("trainId")
		var train models.Train
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(trainId)

		//validate the request body
		if err := c.BindJSON(&train); err != nil {
			c.JSON(http.StatusBadRequest, responses.TrainResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&train); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.TrainResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"name": train.Name, "source": train.Source, "destination": train.Destination, "arrivalTime": train.ArrivalTime, "departureTime": train.DepartureTime}
		result, err := trainCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated user details
		var updatedTrain models.Train
		if result.MatchedCount == 1 {
			err := trainCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedTrain)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedTrain}})
	}

}

func DeleteATrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		trainId := c.Param("trainId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(trainId)

		result, err := trainCollection.DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.TrainResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}},
		)
	}
}

func GetAllTrains() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var trains []models.Train
		defer cancel()

		results, err := trainCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleTrain models.Train
			if err = results.Decode(&singleTrain); err != nil {
				c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			trains = append(trains, singleTrain)
		}

		c.JSON(http.StatusOK,
			responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": trains}},
		)
	}
}

func SearchTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var searchTrain models.SearchTrain
		var train models.Train
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&searchTrain); err != nil {
			c.JSON(http.StatusBadRequest, responses.TrainResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&searchTrain); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.TrainResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		err := trainCollection.FindOne(ctx, bson.M{"source": searchTrain.Source, "destination": searchTrain.Destination}).Decode(&train)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": train}})
	}
}

func UpdateSeats() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		reqMap := c.Request.URL.Query()
		trainNumber := reqMap["trainNumber"][0]
		numberOfSeats, _ := strconv.Atoi(reqMap["numberOfSeats"][0])

		var train models.Train
		err := trainCollection.FindOne(ctx, bson.M{"trainnumber": trainNumber}).Decode(&train)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		if len(train.ReservedSeats) < numberOfSeats {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		var reservedSeatMap map[int]bool

		for i := 1; i <= len(train.ReservedSeats); i++ {
			reservedSeatMap[train.ReservedSeats[i]] = true
		}

		for i := 1; i <= len(train.ReservedSeats); i++ {
			_, exist := reservedSeatMap[i]
			if exist == false {
				train.ReservedSeats = append(train.ReservedSeats, i)
			}
		}

		update := bson.M{"reservedseats": train.ReservedSeats}
		_, err = trainCollection.UpdateOne(ctx, bson.M{"trainnumber": trainNumber}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
		fmt.Println("success true!!!")
		c.JSON(http.StatusOK, gin.H{"error": nil, "status": http.StatusOK})
	}
}
