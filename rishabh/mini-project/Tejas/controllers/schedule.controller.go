package controllers

import (
	"context"
	"net/http"
	"tejas/configs"
	"tejas/dto"
	"tejas/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var scheduleCollection *mongo.Collection = configs.GetCollection(configs.DB, "schedules")

func AddTrainSchedule() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()

		var singleSchedule dto.AddTrainScheduleDTO
		c.BindJSON(&singleSchedule)

		var schedule models.Schedule
		err := scheduleCollection.FindOne(ctx, bson.M{"date": singleSchedule.Date}).Decode(&schedule)

		if err != nil && err.Error() == "mongo: no documents in result" {
			schedule.Date = singleSchedule.Date
			fillTrainDetailsAndAppend(&schedule, singleSchedule)
			_, err := scheduleCollection.InsertOne(ctx, schedule)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Schedule added successfully"})
			return
		} else {
			schedule.Date = singleSchedule.Date
			fillTrainDetailsAndAppend(&schedule, singleSchedule)

			_, err := scheduleCollection.UpdateOne(ctx, bson.M{"date": singleSchedule.Date}, bson.M{"$set": schedule})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Schedule updated successfully"})
			return
		}
	}
}

func fillTrainDetailsAndAppend(schedule *models.Schedule, singleSchedule dto.AddTrainScheduleDTO) {
	var train models.TrainsWithSchedule
	train.Id = singleSchedule.Train.Id
	train.Seats = make([][]bool, singleSchedule.TotalSeats)
	for i := range train.Seats {
		train.Seats[i] = make([]bool, len(singleSchedule.Train.Stations))
	}
	train.Stations = singleSchedule.Train.Stations
	train.PerStationCharge = singleSchedule.Train.PerStationCharge
	if train.PerStationCharge == 0 {
		train.PerStationCharge = 500
	}
	schedule.Trains = append(schedule.Trains, train)
}
