package controllers

import (
	"context"
	"net/http"
	"sanitaria-microservices/appointmentModule/configs"
	"sanitaria-microservices/appointmentModule/models"
	"sanitaria-microservices/appointmentModule/responses"
	"sanitaria-microservices/appointmentModule/services"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var appointmentCollection *mongo.Collection = configs.GetCollection(configs.DB, "appointments")

const (
    topic         = "Booked-appointment"
)

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

func BookAppointment() gin.HandlerFunc{
	return func (c *gin.Context)  {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Param("userId")
		userId, _ := primitive.ObjectIDFromHex(id)
		defer cancel()

		results, err := appointmentCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		var appointment models.Appointment
		appointmentSuccessfull := false
		for results.Next(ctx) {
			if err = results.Decode(&appointment); err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}
			if !appointment.Occupied{
				appointment.Occupied = true
				appointment.GeneralUserID = userId
				appointmentSuccessfull = true
				break
			}
			
		}
		
		if appointmentSuccessfull {
			services.DeleteAppointmentFromDB(appointment)

			p, err_ :=  services.CreateProducer()
			if err_ != nil{
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err_.Error()}})
				return
			}
			services.ProduceAppointment(p,topic,appointment)
			//services.ProduceAppointment(p,topic,appointment)
			c.JSON(http.StatusOK,
				   responses.UserResponse{Status: http.StatusOK, Message: "success",
				   Data: map[string]interface{}{"data": appointment}},
			)
		}else{
			c.JSON(http.StatusOK,
				responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "No appointment available currently."}},
			)
		}
		
	}
		
}
