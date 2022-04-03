package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AvailTicket struct {
	Id             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Train_id       primitive.ObjectID `json:"trainid,omitempty"`
	Capacity       int                `json:"capacity"`
	Price          int                `json:"price"`
	Departure      string             `json:"departure"`
	Arrival        string             `json:"arrival"`
	Departure_time string             `json:"departure_time"`
	Arrival_time   string             `json:"arrival_time"`
}
