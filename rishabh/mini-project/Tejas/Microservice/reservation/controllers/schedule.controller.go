package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reservationService/configs"
	"reservationService/dto"
	"reservationService/kafka"
	"reservationService/models"
	"reservationService/services"
	"reservationService/utils"

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

func ReserveSeat() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()

		var rs dto.ReserveSeatDTO
		c.BindJSON(&rs)

		fmt.Println(rs.Date.UTC())
		var schedule models.Schedule
		err := models.ScheduleCollection.FindOne(ctx, bson.M{"date": rs.Date}).Decode(&schedule)

		if err != nil && err.Error() == "mongo: no documents in result" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			var train models.TrainsWithSchedule
			for _, t := range schedule.Trains {
				if t.Id == rs.TrainId {
					train = t
					break
				}
			}

			if train.Id != rs.TrainId {
				c.JSON(http.StatusNotFound, gin.H{"error": "Train not found"})
			}

			var fromIndex = -1
			var toIndex = -1

			for i, station := range train.Stations {
				if station.Code == rs.From {
					fromIndex = i
				}
				if station.Code == rs.To {
					toIndex = i
				}
			}
			if fromIndex < toIndex && fromIndex != -1 && toIndex != -1 {
				for seatNumber, seat := range train.Seats {
					available := true
					for i := fromIndex; i <= toIndex; i++ {
						if seat[i] {
							available = false
							break
						}
					}
					if available {
						for i := fromIndex; i <= toIndex; i++ {
							seat[i] = true
						}
						_, err := models.ScheduleCollection.UpdateOne(ctx, bson.M{"date": rs.Date}, bson.M{"$set": schedule})
						if err != nil {
							c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
							return
						}
						user, _ := c.MustGet("user_details").(services.SignedDetails)

						amount := train.PerStationCharge * (toIndex - fromIndex + 1)
						paymentDetails, err := payment(c, amount)
						if err != nil {
							c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
							return
						}
						ticket, err := TicketGeneration(user.UserId, rs.TrainId, rs.From, rs.To, rs.Date, paymentDetails.TransactionId, seatNumber)
						if err != nil {
							c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
							return
						}
						go kafka.TicketDetails(ticket)
						c.JSON(http.StatusOK, gin.H{"ticket": ticket})
						return
					}
				}

			}
			c.JSON(http.StatusNotFound, gin.H{"error": "Seat not found"})
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

type paymentDetails struct {
	TransactionId string `json:"transaction_id"`
	UserId        string `json:"user_id"`
	Amount        int    `json:"amount"`
	Status        string `json:"status"`
}
type paymentResponse struct {
	Message string `json:"message"`
	Data    paymentDetails
}

func payment(c *gin.Context, amount int) (paymentDetails, error) {
	var pd paymentDetails

	data := map[string]int{}
	data["amount"] = amount

	jsonData, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", configs.PaymentGatewayUrl(), bytes.NewBuffer(jsonData))
	if err != nil {
		return pd, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.GetHeader("Authorization"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return pd, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return pd, err
	}

	var response paymentResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return pd, err
	}

	pd = response.Data

	return pd, nil
}
