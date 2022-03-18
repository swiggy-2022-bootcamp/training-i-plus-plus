package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AvailTicket struct {
	Id             primitive.ObjectID `json:"availTicketId,omitempty"`
	Train_id       primitive.ObjectID `json:"TrainId,omitempty"`
	Capacity       int                `json:"capacity"`
	Price          int                `json:"price"`
	Departure      string             `json:"departure"`
	Arrival        string             `json:"arrival"`
	Departure_time time.Time          `json:"departure_time"`
	Arrival_time   time.Time          `json:"arrival_time"`
}
