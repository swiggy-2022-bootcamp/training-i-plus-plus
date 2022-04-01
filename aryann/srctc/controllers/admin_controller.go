package controllers

import (
	"context"
	"fmt"
	"net/http"
	"srctc/kafka"
	"srctc/models"
	"srctc/responses"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var avalidate = validator.New()

const layout = "Jan 2, 2006 at 3:04pm (MST)"

func init() {
	go kafka.Consume_booked_ticket_for_avail()
}

func GetAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		adminId := c.Param("adminid")
		var admin models.Admin
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(adminId)

		admin, err := adminRepo.Read(objId)
		// err := adminCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&admin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": admin}})
	}
}

func DeleteAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		adminId := c.Param("adminid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(adminId)

		result, err := adminRepo.Delete(objId)
		// result, err := adminCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// if result.(int) < 1 {
		// 	c.JSON(http.StatusNotFound,
		// 		responses.AdminResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Admin with specified ID not found!"}},
		// 	)
		// 	return
		// }

		c.JSON(http.StatusOK,
			responses.AdminResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Admin successfully deleted!", "result": result}},
		)
	}
}

func CreateTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var train models.Train
		defer cancel()

		if err := c.BindJSON(&train); err != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := avalidate.Struct(&train); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AdminResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newTrain := models.Train{
			Source:      train.Source,
			Destination: train.Destination,
		}

		result, err := trainRepo.Create(newTrain)
		// result, err := trainCollection.InsertOne(ctx, newTrain)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AdminResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		go kafka.Produce_train(newTrain)

		c.JSON(http.StatusCreated, responses.TrainResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		trainId := c.Param("trainid")
		var train models.Train
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(trainId)

		train, err := trainRepo.Read(objId)
		// err := trainCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&train)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": train}})
	}
}

func EditTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		trainId := c.Param("trainid")
		var train models.Train
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(trainId)

		if err := c.BindJSON(&train); err != nil {
			c.JSON(http.StatusBadRequest, responses.TrainResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := avalidate.Struct(&train); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.TrainResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		// update := bson.M{"station1": train.Source, "station2": train.Destination}
		result, err := trainRepo.Update(train, objId)
		// result, err := trainCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// var updatedTrain models.Train
		// if result.(int) == 1 {
		// 	err := trainCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedTrain)
		// 	if err != nil {
		// 		c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		// 		return
		// 	}
		// }

		c.JSON(http.StatusOK, responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func DeleteTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		trainId := c.Param("trainid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(trainId)

		result, err := trainRepo.Delete(objId)
		fmt.Println(result)
		// result, err := trainCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			fmt.Println("iam here")
			c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// if result.(int) < 1 {
		// 	c.JSON(http.StatusInternalServerError, responses.TrainResponse{Status: http.StatusInternalServerError, Message: "error id not found", Data: map[string]interface{}{"data": err.Error()}})
		// 	return
		// }

		c.JSON(http.StatusOK, responses.TrainResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Train successfully deleted!"}})
	}
}

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

		go kafka.Produce_avail_ticket(newTicket)

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

func GetAllTickets() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// var tickets []models.Ticket
		defer cancel()

		results, err := ticketRepo.ReadAll()
		// results, err := ticketCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.TicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// defer results.Close(ctx)
		// for results.Next(ctx) {
		// 	var singleTicket models.Ticket
		// 	if err = results.Decode(&singleTicket); err != nil {
		// 		c.JSON(http.StatusInternalServerError, responses.TicketResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		// 	}

		// 	tickets = append(tickets, singleTicket)
		// }

		c.JSON(http.StatusOK,
			responses.TicketResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": results}},
		)
	}
}
