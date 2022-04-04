package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"ticket_service/configs"
	"ticket_service/kafka"
	"ticket_service/models"
	"ticket_service/responses"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ticketCollection *mongo.Collection = configs.GetCollection(configs.DB, "ticket_details")
var validate = validator.New()
var trainURL = "http://localhost:6000/trains/"

func GetAllTickets() gin.HandlerFunc {
	return func(c *gin.Context) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			var tickets []models.Ticket
			defer cancel()

			user_details, _ := c.Get("user_details")
			fmt.Print("hello",user_details)
			
			var query bson.M;
			if user_details.(jwt.MapClaims)["role"] == "ADMIN" {
				query = bson.M{}
			} else {
				query =  bson.M{"id": user_details.(jwt.MapClaims)["userId"]}
			}
			results, err := ticketCollection.Find(ctx, query)

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
				
					tickets = append(tickets, singleTicket)
			}

			c.JSON(http.StatusOK,
					responses.TicketResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": tickets}},
			)
	}
}

type logWriter struct{}

func GetTicketById() gin.HandlerFunc {
	return func(c *gin.Context) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			ticketId := c.Param("ticketId")
			var ticket models.Ticket
			defer cancel()

			objId, _ := primitive.ObjectIDFromHex(ticketId)

			err := ticketCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&ticket)
			if err != nil {
					c.JSON(http.StatusInternalServerError, responses.TicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
					return
			}

			c.JSON(http.StatusOK, responses.TicketResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": ticket}})
	}
}

func BookTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			var ticket models.Ticket
			defer cancel()

			//validate the request body
			if err := c.BindJSON(&ticket); err != nil {
					c.JSON(http.StatusBadRequest, responses.TicketResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
					return
			}

			//use the validator library to validate required fields
			if validationErr := validate.Struct(&ticket); validationErr != nil {
					c.JSON(http.StatusBadRequest, responses.TicketResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
					return
			}

			//Getting Train details from Train service
			client := &http.Client{
        Timeout: time.Second * 10,
    	}
			req, err := http.NewRequest("GET", trainURL+string(ticket.TrainNumber), nil)
			if err != nil {
				c.JSON(http.StatusBadRequest, responses.TicketResponse{Status: http.StatusBadRequest, Message: "Train Not Found", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			req.Header.Set("Authorization", c.Request.Header.Get("Authorization"))
			response, err := client.Do(req)
			if err != nil {
				c.JSON(http.StatusBadRequest, responses.TicketResponse{Status: http.StatusBadRequest, Message: "Train Not Found", Data: map[string]interface{}{"data": err.Error()}})
			  return
			}
			defer response.Body.Close()

			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				c.JSON(http.StatusBadRequest, responses.TicketResponse{Status: http.StatusBadRequest, Message: "Train Not Found", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		
			trainResponse := responses.TicketResponse{}
			
			jsonErr := json.Unmarshal(body, &trainResponse)
			if jsonErr != nil {
				c.JSON(http.StatusBadRequest, responses.TicketResponse{Status: http.StatusBadRequest, Message: "Train Not Found", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			if trainResponse.Status != 200 {
				c.JSON(http.StatusBadRequest, responses.TicketResponse{Status: http.StatusBadRequest, Message: "Internal Error while getting train details", Data: map[string]interface{}{"data": trainResponse.Message}})
				return
			} else {
				trainData := trainResponse.Data["data"]
				if trainData, ok := trainData.(map[string]interface{}); ok {
					if trainData["availableseats"].(float64) < float64(ticket.NoOfSeats) {
						c.JSON(http.StatusBadRequest, responses.TicketResponse{Status: http.StatusBadRequest, Message: "Seats not avialable to book"})
						return
					} else {
						newTicket := models.Ticket{
							Id:          primitive.NewObjectID(),
							TrainNumber: ticket.TrainNumber,
							NoOfSeats: 	 ticket.NoOfSeats,
							UserName:    ticket.UserName,
							Passengers : ticket.Passengers,   
							TotalCost:   float64(ticket.NoOfSeats) * trainData["ticketprice"].(float64),
							Status   :   "Booked",
						}
							
						result, err := ticketCollection.InsertOne(ctx, newTicket)
						if err != nil {
							c.JSON(http.StatusInternalServerError, responses.TicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
							return
						}
						go kafka.UpdateSeatsProducer(ticket.TrainNumber, ticket.NoOfSeats, true)
			      c.JSON(http.StatusCreated, responses.TicketResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
					}
				}
			}
	}
}


func CancelTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			ticketId := c.Param("ticketId")
			defer cancel()
		
			objId, _ := primitive.ObjectIDFromHex(ticketId)
		
			update := bson.M{
				"status": "Canceled",
			}
			result, err := ticketCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
		
			if err != nil {
					c.JSON(http.StatusInternalServerError, responses.TicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
					return
			}
		
			//get updated ticket details
			var updateticket models.Ticket
			if result.MatchedCount == 1 {
					err := ticketCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updateticket)
					if err != nil {
							c.JSON(http.StatusInternalServerError, responses.TicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
							return
					}
			}
			go kafka.UpdateSeatsProducer(updateticket.TrainNumber, updateticket.NoOfSeats, false)
			c.JSON(http.StatusOK, responses.TicketResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updateticket}})
	}
}