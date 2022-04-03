package controllers

import (
	"context"
	"net/http"
	"srctc/kafka"
	"srctc/models"
	"srctc/repository"
	"srctc/responses"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

var purchasedRepo repository.PurchasedRepository
var ticketRepo repository.TicketRepository
var trainRepo repository.TrainRepository

func init() {
	go kafka.Consume_ticket()
	go kafka.Consume_train()
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		user, err := userRepo.Read(objId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		result, err := userRepo.Delete(objId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		signup, err := userRepo.Delete(objId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!", "result": result, "signup": signup}},
		)
	}
}

func PurchaseTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var purchased models.Purchased
		defer cancel()

		if err := c.BindJSON(&purchased); err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error in binding", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&purchased); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error in validating", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		var ticket models.Ticket

		ticket, err := ticketRepo.ReadTrainId(purchased.Train_id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PurchasedResponse{Status: http.StatusInternalServerError, Message: "Incorrect train id", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if ticket.Capacity == 0 {
			c.JSON(http.StatusInternalServerError, responses.PurchasedResponse{Status: http.StatusInternalServerError, Message: "No tickets available", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		ticket.Capacity = ticket.Capacity - 1
		_, err = ticketRepo.Update(ticket, ticket.ID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TicketResponse{Status: http.StatusInternalServerError, Message: "error in updating capacity", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var trainbooked models.Train
		trainbooked, err = trainRepo.Read(purchased.Train_id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PurchasedResponse{Status: http.StatusInternalServerError, Message: "error in train find", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		newpurchased := models.Purchased{
			Train_id:       purchased.Train_id,
			User_id:        purchased.User_id,
			Departure:      trainbooked.Source,
			Arrival:        trainbooked.Destination,
			Departure_time: ticket.Departure_time,
			Arrival_time:   ticket.Arrival_time,
			Cost:           ticket.Cost,
		}

		go kafka.Produce_purchased_ticket(newpurchased)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.PurchasedResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": "Ticket successfully purchased!"}})
	}
}

func GetPurchased() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		purchasedId := c.Param("purchasedid")
		var purchased models.Purchased
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(purchasedId)

		purchased, err := purchasedRepo.Read(objId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PurchasedResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.PurchasedResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": purchased}})
	}
}

func DeletePurchased() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		purchasedId := c.Param("purchasedid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(purchasedId)

		purchased, err := purchasedRepo.Read(objId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PurchasedResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		result, err := purchasedRepo.Delete(objId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PurchasedResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK,
			responses.PurchasedResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Purchased successfully deleted!", "result": result, "purchased": purchased}},
		)
	}
}
