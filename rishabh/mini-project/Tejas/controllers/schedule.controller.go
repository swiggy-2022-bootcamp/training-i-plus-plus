package controllers

import (
	"context"
	"fmt"
	"net/http"
	"tejas/dto"
	"tejas/models"
	"tejas/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func AddTrainSchedule() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()

		var singleSchedule dto.AddTrainScheduleDTO
		c.BindJSON(&singleSchedule)

		var schedule models.Schedule
		err := models.ScheduleCollection.FindOne(ctx, bson.M{"date": singleSchedule.Date}).Decode(&schedule)

		if err != nil && err.Error() == "mongo: no documents in result" {
			schedule.Date = singleSchedule.Date
			fillTrainDetailsAndAppend(&schedule, singleSchedule)
			_, err := models.ScheduleCollection.InsertOne(ctx, schedule)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Schedule added successfully"})
			return
		} else {
			schedule.Date = singleSchedule.Date
			fillTrainDetailsAndAppend(&schedule, singleSchedule)

			_, err := models.ScheduleCollection.UpdateOne(ctx, bson.M{"date": singleSchedule.Date}, bson.M{"$set": schedule})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Schedule updated successfully"})
			return
		}
	}
}

func Availabilty() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()

		var fromStation = c.Query("from")
		var toStation = c.Query("to")
		date, err := utils.ISODateToTime(c.Query("date"))

		if err != nil {
			fmt.Println(err)
		}

		var schedule models.Schedule
		err = models.ScheduleCollection.FindOne(ctx, bson.M{"date": date}).Decode(&schedule)

		if err != nil && err.Error() == "mongo: no documents in result" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			var trains []models.TrainsWithSchedule
			for _, train := range schedule.Trains {
				var fromIndex = -1
				var toIndex = -1

				for i, station := range train.Stations {
					if station.Code == fromStation {
						fromIndex = i
					}
					if station.Code == toStation {
						toIndex = i
					}
				}
				if fromIndex < toIndex && fromIndex != -1 && toIndex != -1 {
					for _, seat := range train.Seats {
						available := true
						for i := fromIndex; i <= toIndex; i++ {
							if seat[i] {
								available = false
								break
							}
						}
						if available {
							trains = append(trains, train)
							break
						}
					}
				}
			}
			trainsResponse := make([]models.TrainWithoutSeats, len(trains))
			for i, train := range trains {
				trainsResponse[i].Id = train.Id
				trainsResponse[i].Stations = train.Stations
				trainsResponse[i].PerStationCharge = train.PerStationCharge
			}
			c.JSON(http.StatusOK, gin.H{"trains": trainsResponse})
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
