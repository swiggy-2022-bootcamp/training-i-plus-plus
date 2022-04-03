package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/taran1515/crud/configs"
	"github.com/taran1515/crud/models"
	"github.com/taran1515/crud/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"net/http"
	"time"
)

var ticketCollection = configs.GetCollection(configs.DB, "ticket")
var validate *validator.Validate = validator.New()

func BookTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
		var ticket models.Ticket
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&ticket); err != nil {
			c.JSON(http.StatusBadRequest, responses.TrainResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&ticket); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.TrainResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		ticketStatus := models.TicketStatus(
			models.Reserved)

		newTicket := models.Ticket{
			TicketPnr:     primitive.NewObjectID(),
			PassengerName: ticket.PassengerName,
			Source:        ticket.Source,
			Destination:   ticket.Destination,
			Amount:        ticket.Amount,
			SeatNumbers:   ticket.SeatNumbers,
			Distance:      ticket.Distance,
			Quota:         ticket.Quota,
			TrainNumber:   ticket.TrainNumber,
			TicketStatus:  ticketStatus,
		}

		result, err := ticketCollection.InsertOne(ctx, newTicket)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.TrainResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetATicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		trainId := c.Param("trainId")
		var ticket models.Ticket
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(trainId)

		err := ticketCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&ticket)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": ticket}})
	}
}

func CancelBooking() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		ticketId := c.Param("ticketId")
		var ticket models.Ticket
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(ticketId)

		//validate the request body
		if err := c.BindJSON(&ticket); err != nil {
			c.JSON(http.StatusBadRequest, responses.TrainResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&ticket); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.TrainResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		ticketStatus := models.TicketStatus(
			models.Reserved)

		update := bson.M{"TicketStatus": ticketStatus}
		result, err := ticketCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated user details
		var updatedTrain models.Ticket
		if result.MatchedCount == 1 {
			err := ticketCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedTrain)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedTrain}})
	}

}
