package services

import (
	"context"
	"fmt"
	"sanitaria-microservices/appointmentModule/models"
	"sanitaria-microservices/appointmentModule/configs"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var appointmentCollection *mongo.Collection = configs.GetCollection(configs.DB, "appointments")

func AddAppointmentToDB(appointment models.Appointment){
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	_, err := appointmentCollection.InsertOne(ctx, appointment)
    if err != nil {
        fmt.Println(err.Error())
    }
}

func DeleteAppointmentFromDB(appointment models.Appointment){
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	_, err := appointmentCollection.DeleteOne(ctx, bson.M{"_id":appointment.Id})
    if err != nil {
        fmt.Println(err.Error())
    }
}

