package controllers

import (
	"context"
	"net/http"
	"srctc/database"
	"srctc/models"
	"srctc/responses"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var adminCollection *mongo.Collection = database.GetCollection(database.DB, "admins")
var trainCollection *mongo.Collection = database.GetCollection(database.DB, "trains")
var availticketCollection *mongo.Collection = database.GetCollection(database.DB, "tickets")
var avalidate = validator.New()

const layout = "Jan 2, 2006 at 3:04pm (MST)"

func CreateAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var admin models.Admin
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&admin); err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&admin); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newAdmin := models.Admin{
			Name:  admin.Name,
			Email: admin.Email,
		}

		result, err := adminCollection.InsertOne(ctx, newAdmin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.AdminResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		adminId := c.Param("adminid")
		var admin models.Admin
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(adminId)

		err := adminCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&admin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": admin}})
	}
}

func DeleteAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		adminId := c.Param("adminid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(adminId)

		result, err := adminCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.AdminResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Admin with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Admin successfully deleted!"}},
		)
	}
}

func CreateTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var train models.Train
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&train); err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&train); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newTrain := models.Train{
			Source:      train.Source,
			Destination: train.Destination,
		}

		result, err := trainCollection.InsertOne(ctx, newTrain)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.TrainResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		trainId := c.Param("trainid")
		var train models.Train
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(trainId)

		err := trainCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&train)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": train}})
	}
}

func EditTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		trainId := c.Param("trainid")
		var train models.Train
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(trainId)

		//validate the request body
		if err := c.BindJSON(&train); err != nil {
			c.JSON(http.StatusBadRequest, responses.TrainResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&train); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.TrainResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"station1": train.Source, "station2": train.Destination}
		result, err := trainCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated train details
		var updatedTrain models.Train
		if result.MatchedCount == 1 {
			err := trainCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedTrain)
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
		trainId := c.Param("trainid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(trainId)

		result, err := trainCollection.DeleteOne(ctx, bson.M{"_id": objId})

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

func CreateTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var availticket models.Ticket
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&availticket); err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&availticket); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		_, err1 := time.Parse(layout, availticket.Departure_time)
		_, err := time.Parse(layout, availticket.Arrival_time)

		if err != nil || err1 != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "time not in correct format", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		newTicket := models.Ticket{
			Train_id:       availticket.Train_id,
			Capacity:       availticket.Capacity,
			Cost:           availticket.Cost,
			Departure:      availticket.Departure,
			Arrival:        availticket.Arrival,
			Departure_time: availticket.Departure_time,
			Arrival_time:   availticket.Arrival_time,
		}

		result, err := availticketCollection.InsertOne(ctx, newTicket)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.TicketResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		availticketId := c.Param("availticketid")
		var availticket models.Ticket
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(availticketId)

		err := availticketCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&availticket)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.TicketResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": availticket}})
	}
}

func DeleteTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		availticketId := c.Param("availticketid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(availticketId)

		result, err := availticketCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.TicketResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Ticket with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.TicketResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Ticket successfully deleted!"}},
		)
	}
}

func GetAllTickets() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var availtickets []models.Ticket
		defer cancel()

		results, err := availticketCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleTicket models.Ticket
			if err = results.Decode(&singleTicket); err != nil {
				c.JSON(http.StatusInternalServerError, responses.TicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			availtickets = append(availtickets, singleTicket)
		}

		c.JSON(http.StatusOK,
			responses.TicketResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": availtickets}},
		)
	}
}
