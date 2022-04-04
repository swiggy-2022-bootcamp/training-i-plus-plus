package service

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"train/database"
	"train/helper"
	"train/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var trainCollection *mongo.Collection = database.GetCollection(database.DB, "trains")
var validate *validator.Validate = validator.New()

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

func CheckSeatAvailablity() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		//api request from postman
		// var checkAvailablityReq model.CheckAvailabiltyRequest
		// if err := c.BindJSON(&checkAvailablityReq); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
		// 	return
		// }
		// fmt.Println(checkAvailablityReq)
		// var train model.Train
		// if err := trainCollection.FindOne(ctx, bson.M{"trainnumber": checkAvailablityReq.TrainNumber}).Decode(&train); err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusInternalServerError, "success": false})
		// 	return
		// }
		// fmt.Println("train service, availablity ", train)
		// if train.AvailableSeatCount < checkAvailablityReq.NumOfSeats {
		// 	c.JSON(http.StatusBadGateway, gin.H{"error": "tickets are not available", "success": false})
		// 	return
		// }
		// fmt.Println("seats are availavble")
		// train.AvailableSeatCount -= checkAvailablityReq.NumOfSeats
		// _, err := trainCollection.UpdateOne(ctx, bson.M{"trainnumber": checkAvailablityReq.TrainNumber}, bson.M{"$set": train})
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		// 	return
		// }

		reqMap := c.Request.URL.Query()
		trainnumber := reqMap["trainnumber"][0]
		numofseats, _ := strconv.Atoi(reqMap["numofseats"][0])
		incrementcount, _ := strconv.ParseBool(reqMap["incrementcount"][0])

		fmt.Println("params ", trainnumber, numofseats, incrementcount)
		fmt.Printf("%T  %T", trainnumber, numofseats, incrementcount)

		var train model.Train
		if err := trainCollection.FindOne(ctx, bson.M{"trainnumber": trainnumber}).Decode(&train); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusInternalServerError, "success": false})
			return
		}
		fmt.Println("train service, availablity ", train)
		if train.AvailableSeatCount < numofseats {
			c.JSON(http.StatusBadGateway, gin.H{"error": "tickets are not available", "success": false})
			return
		}
		fmt.Println("seats are availavble")
		if incrementcount {
			train.AvailableSeatCount += numofseats
		} else {
			train.AvailableSeatCount -= numofseats
		}
		_, err := trainCollection.UpdateOne(ctx, bson.M{"trainnumber": trainnumber}, bson.M{"$set": train})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
		fmt.Println("success true!!!")
		c.JSON(http.StatusOK, gin.H{"error": nil, "status": http.StatusOK})
	}
}

// ShowAccount godoc
// @Summary      Search trains
// @Description  Finding trains running between departure station and arrival station
// @Tags         search
// @Accept       json
// @Produce      json
// @Param        DepartureStation  			body 	string  true  "station name from where train starts"
// @Param        ArrivalStation 		body	string   	true  "station name where train journey ends"
// @Param        DepartureDate 		body	string  true  "data of departure"
// @Success      200  {object}  model.SearchTrainResponse
// @Failure      400  {number} 	http.StatusBadRequest
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router       /search_trains [get]
func SearchTrains() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var searchReq model.SearchTrainRequest
		var trains []model.SearchTrainResponse

		if err := c.BindJSON(&searchReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
			return
		}
		token := c.Request.Header.Get("Cookie")
		claims, msg := helper.ValidateToken(token)
		fmt.Println(claims, msg)

		results, err := trainCollection.Find(ctx, bson.M{"departurestation": searchReq.DepartureStation,
			"arrivalstation": searchReq.ArrivalStation})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
			return
		}
		defer results.Close(ctx)
		fmt.Println(searchReq.DepartureStation, searchReq.ArrivalStation, results)
		for results.Next(ctx) {
			var train model.Train
			if err = results.Decode(&train); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
				return
			}
			if train.AvailableSeatCount > 0 {
				singleSearchResp := model.SearchTrainResponse{
					TrainNumber:        train.TrainNumber,
					TrainName:          train.TrainName,
					AvailableSeatCount: train.AvailableSeatCount,
					Fare:               train.Fare,
				}
				trains = append(trains, singleSearchResp)
			}
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "success", "data": trains})

	}
}
