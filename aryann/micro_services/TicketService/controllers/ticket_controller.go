package controllers

import (
	"TicketService/kafka"
	"TicketService/middlewares"
	"TicketService/models"
	"TicketService/repository"
	"TicketService/responses"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	avalidate  = validator.New()
	ticketRepo repository.TicketRepository
)

func init() {
	go kafka.Consume_purchased_ticket()
}

const layout = "Jan 2, 2006 at 3:04pm (MST)"

func CreateTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var ticket models.Ticket
		defer cancel()

		if err := c.BindJSON(&ticket); err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := avalidate.Struct(&ticket); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		_, err1 := time.Parse(layout, ticket.Departure_time)
		_, err := time.Parse(layout, ticket.Arrival_time)

		if err != nil || err1 != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "time not in correct format", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		newTicket := models.Ticket{
			Train_id:       ticket.Train_id,
			Capacity:       ticket.Capacity,
			Cost:           ticket.Cost,
			Departure_time: ticket.Departure_time,
			Arrival_time:   ticket.Arrival_time,
		}

		result, err := ticketRepo.Create(newTicket)
		// result, err := ticketCollection.InsertOne(ctx, newTicket)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// go kafka.Produce_ticket(newTicket)

		c.JSON(http.StatusCreated, responses.TicketResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		ticketId := c.Param("ticketid")
		// var ticket models.Ticket
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(ticketId)

		result, err := ticketRepo.Read(objId)
		// err := ticketCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&ticket)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.TicketResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func DeleteTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		ticketId := c.Param("ticketid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(ticketId)

		result, err := ticketRepo.Delete(objId)
		// result, err := ticketCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// if result.(int) < 1 {
		// 	c.JSON(http.StatusNotFound,
		// 		responses.TicketResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Ticket with specified ID not found!"}},
		// 	)
		// 	return
		// }

		c.JSON(http.StatusOK,
			responses.TicketResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Ticket successfully deleted!", "result": result}},
		)
	}
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func IsAuthorized(group string) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.Request.Header.Get("Authorization")

		strArr := strings.Split(bearToken, " ")
		if len(strArr) != 2 {
			respondWithError(c, 401, "No bearer token")
			return
		}

		token, err := jwt.Parse(strArr[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Fatal("Unexpected signing method")
				return nil, fmt.Errorf(("invalid signing method"))
			}

			return middlewares.GetMySigingKey(), nil
		})
		if err != nil {
			respondWithError(c, 501, err.Error())
			return
		}
		if !token.Valid {
			respondWithError(c, 401, "Invalid token")
			return
		}

		if token.Claims.(jwt.MapClaims)["group"] != group {
			respondWithError(c, 401, "unauthorized user")
			return
		}

		c.Next()
	}
}
