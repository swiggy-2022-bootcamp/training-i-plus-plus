package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"train_service/configs"
	"train_service/kafka"
	"train_service/models"
	"train_service/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var trainCollection *mongo.Collection = configs.GetCollection(configs.DB, "train_details")
var validate = validator.New()

func init() {
	go kafka.UpdateSeatsConsumer()
}

func AddTrain() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var train models.Train
        defer cancel()

				const layout = "2006-01-02T15:04:05.000Z"
				train.ArrivalTime, _ = time.Parse(layout, train.ArrivalTime.String())
				train.DepartureTime, _ = time.Parse(layout, train.DepartureTime.String())
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

        newTrain := models.Train{
            Id:       primitive.NewObjectID(),
            Name:     train.Name,
            Source: 	train.Source,
            Destination: train.Destination,
						TotalSeats : train.TotalSeats,   
						AvailableSeats: train.AvailableSeats,
						ArrivalTime   : train.ArrivalTime,
						DepartureTime : train.DepartureTime,
						TicketPrice   : train.TicketPrice,
        }
      
        result, err := trainCollection.InsertOne(ctx, newTrain)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusCreated, responses.TrainResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
    }
}


func GetTrainById() gin.HandlerFunc {
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


func UpdateTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			trainId := c.Param("trainId")
			var train models.Train
			defer cancel()
			objId, _ := primitive.ObjectIDFromHex(trainId)


			const layout = "2006-01-02T15:04:05.000Z"
			train.ArrivalTime, _ = time.Parse(layout, train.ArrivalTime.String())
			train.DepartureTime, _ = time.Parse(layout, train.DepartureTime.String())
			
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

			update := bson.M{
				"name": train.Name, 
				"source": train.Source,
				"destination": train.Destination,
				"totalseats": train.TotalSeats,
				"availableseats":train.AvailableSeats,
				"arrivaltime": train.ArrivalTime,
				"departuretime":train.DepartureTime,
				"ticketprice": train.TicketPrice,
			}
			result, err := trainCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
			if err != nil {
					c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
					return
			}

			//get updated train details
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

func DeleteTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			trainId := c.Param("trainId")
			fmt.Print((trainId))
			defer cancel()

			objId, _ := primitive.ObjectIDFromHex(trainId)
			fmt.Print((objId))

			result, err := trainCollection.DeleteOne(ctx, bson.M{"id": objId})
			if err != nil {
					c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
					return
			}

			if result.DeletedCount < 1 {
					c.JSON(http.StatusNotFound,
							responses.TrainResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Train with specified ID not found!"}},
					)
					return
			}

			c.JSON(http.StatusOK,
					responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Train successfully deleted!"}},
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
