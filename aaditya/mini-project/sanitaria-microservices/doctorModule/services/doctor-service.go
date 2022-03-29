package services

import (
	"context"
	"sanitaria-microservices/doctorModule/models"
	"sanitaria-microservices/doctorModule/configs"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var doctorCollection *mongo.Collection = configs.GetCollection(configs.DB, "doctors")

func UpdateDoctorDB(appointment models.Appointment){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	doctorId := appointment.DoctorID
	
    var doctor models.Doctor
	err := doctorCollection.FindOne(ctx, bson.M{"_id": doctorId}).Decode(&doctor)
        if err != nil {
            return
        }
	appointmentsList := doctor.Appointments

	for ind,entry := range appointmentsList{
		if entry.Id == appointment.Id{
			appointmentsList[ind] = appointment
			break
		}
	}

	update := bson.M{"appointments" : appointmentsList}

	_, err = doctorCollection.UpdateOne(ctx, bson.M{"_id": doctorId}, bson.M{"$set": update})

	if err != nil {
		return
	}
}