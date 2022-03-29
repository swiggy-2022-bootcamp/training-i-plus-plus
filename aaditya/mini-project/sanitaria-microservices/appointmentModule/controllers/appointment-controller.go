package controllers

import (
	"context"
	"fmt"
	"net/http"
	"sanitaria-microservices/appointmentModule/configs"
	"sanitaria-microservices/appointmentModule/models"
	"sanitaria-microservices/appointmentModule/responses"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var appointmentCollection *mongo.Collection = configs.GetCollection(configs.DB, "appointments")

func GetAllAppointments() gin.HandlerFunc{
	return func (c *gin.Context)  {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var appointments []models.Appointment
		defer cancel()

		results, err := appointmentCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var appointment models.Appointment
			if err = results.Decode(&appointment); err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			appointments = append(appointments, appointment)
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": appointments}},
		)
	}
}

func AddAppointmentToDB(appointment models.Appointment){
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	_, err := appointmentCollection.InsertOne(ctx, appointment)
    if err != nil {
        fmt.Println(err.Error())
    }
}

func BookAppointment() gin.HandlerFunc{
	return func (c *gin.Context)  {
		fmt.Println("Called BookAppointment() ")
	}
}