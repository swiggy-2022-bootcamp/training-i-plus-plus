package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ticket struct {
	ID             primitive.ObjectID `json:"id,omitempty"`
	Train_id       primitive.ObjectID `json:"trainid,omitempty"`
	Capacity       int                `json:"capacity"`
	Cost           int                `json:"cost"`
	Departure_time string             `json:"departure_time"`
	Arrival_time   string             `json:"arrival_time"`
}
