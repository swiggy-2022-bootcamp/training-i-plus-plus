package controllers

import (
	//"context"
	//"errors"
	"fmt"
	//"net/http"
	"sanitaria-microservices/appointmentModule/configs"
	// "sanitaria-microservices/appointmentModule/models"
	// "sanitaria-microservices/appointmentModule/responses"
	// "sanitaria-microservices/appointmentModule/services"
	//"time"
	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	 "go.mongodb.org/mongo-driver/mongo"
)

var appointmentCollection *mongo.Collection = configs.GetCollection(configs.DB, "appointments")
// var validate = validator.New()

func GetAllAppointments() gin.HandlerFunc{
	return func (c *gin.Context)  {
		fmt.Println("Called GetAllAppointment() ")
	}
}


func BookAppointment() gin.HandlerFunc{
	return func (c *gin.Context)  {
		fmt.Println("Called BookAppointment() ")
	}
}