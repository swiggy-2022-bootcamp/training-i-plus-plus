package controller

import (
	"context"
	"fmt"
	"net/http"
	"ticket_reservation_system/config"
	"ticket_reservation_system/model"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var trainCollection *mongo.Collection = config.GetCollection(config.DB, "trains")

func AddTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var train model.Train
		defer cancel()
		//validate the request body
		if err := c.BindJSON(&train); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}
		//use the validator library to validate required fields
		if validationErr := validate.Struct(&train); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": validationErr.Error()})
			return
		}

		newTrain := model.Train{
			ID:                 primitive.NewObjectID(),
			TrainNumber:        train.TrainNumber,
			TrainName:          train.TrainName,
			DepartureStation:   train.DepartureStation,
			ArrivalStation:     train.ArrivalStation,
			DepartureDate:      train.DepartureDate,
			TotalSeatCount:     train.TotalSeatCount,
			AvailableSeatCount: train.AvailableSeatCount,
			Fare:               train.Fare,
		}
		result, err := trainCollection.InsertOne(ctx, newTrain)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "success", "data": result})
	}
}

func GetTrainByTrainNumber() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		trainnumber := c.Param("trainnumber")
		fmt.Print("get train", trainnumber)
		var train model.Train
		defer cancel()
		if err := trainCollection.FindOne(ctx, bson.M{"trainnumber": trainnumber}).Decode(&train); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusCreated, "message": "success", "data": train})
	}
}

func GetAllTrains() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var trains []model.Train
		defer cancel()
		results, err := trainCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
			return
		}
		defer results.Close(ctx)
		for results.Next(ctx) {
			var train model.Train
			if err = results.Decode(&train); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
				return
			}
			trains = append(trains, train)
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "success", "data": trains})
	}
}

func UpdateTrainDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		trainnumber := c.Param("trainnumber")
		var train model.Train
		defer cancel()
		//validate the request body
		if err := c.BindJSON(&train); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}
		//use the validator library to validate required fields
		if validationErr := validate.Struct(&train); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": validationErr.Error()})
			return
		}
		update := bson.M{"_id": train.ID, "train_number": train.TrainName, "train_name": train.TrainName,
			"departure_station": train.DepartureStation, "arrival_station": train.ArrivalStation, "departure_date": train.DepartureDate,
			"total_seat_count": train.TotalSeatCount, "available_seat_count": train.AvailableSeatCount, "fare": train.Fare}
		result, err := trainCollection.UpdateOne(ctx, bson.M{"trainnumber": trainnumber}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
			return
		}
		//get updated train details
		var updatedTrain model.Train
		if result.MatchedCount == 1 {
			if err := trainCollection.FindOne(ctx, bson.M{"trainnumber": trainnumber}).Decode(&updatedTrain); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "success", "data": updatedTrain})
	}
}

func DeleteTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		trainNumber := c.Param("train_number")
		defer cancel()
		result, err := trainCollection.DeleteOne(ctx, bson.M{"trainNumber": trainNumber})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
			return
		}
		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "success", "data": "train successfully deleted"})
	}
}
