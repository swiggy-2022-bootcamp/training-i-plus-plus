package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Station struct {
	Code          string             `json:"code" bson:"code"`
	ArrivalTime   primitive.DateTime `json:"arrivalTime" bson:"arrivalTime"`
	DepartureTime primitive.DateTime `json:"departureTime" bson:"departureTime"`
}

type Train struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Stations []Station          `json:"stations" bson:"stations"`
}
