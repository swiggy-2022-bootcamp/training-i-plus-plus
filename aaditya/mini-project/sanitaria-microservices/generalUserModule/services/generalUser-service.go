package services

import (
	"context"
	"sanitaria-microservices/generalUserModule/models"
	"sanitaria-microservices/generalUserModule/configs"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var generalUserCollection *mongo.Collection = configs.GetCollection(configs.DB, "generalUsers")

func UpdateGeneralUserDB(appointment models.Appointment){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	generalUserId := appointment.GeneralUserID
	
    var generalUser models.GeneralUser
	err := generalUserCollection.FindOne(ctx, bson.M{"_id": generalUserId}).Decode(&generalUser)
        if err != nil {
            return
        }
	appointmentsList := generalUser.Appointments

	appointmentsList = append(appointmentsList, appointment)

	update := bson.M{"appointments" : appointmentsList}

	_, err = generalUserCollection.UpdateOne(ctx, bson.M{"_id": generalUserId}, bson.M{"$set": update})

	if err != nil {
		return
	}
}