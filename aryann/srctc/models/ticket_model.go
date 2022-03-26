package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AvailTicket struct {
	Train_id       primitive.ObjectID `json:"trainid,omitempty"`
	Capacity       int                `json:"capacity"`
	Cost           int                `json:"cost"`
	Departure      string             `json:"departure"`
	Arrival        string             `json:"arrival"`
	Departure_time string             `json:"departure_time"`
	Arrival_time   string             `json:"arrival_time"`
}
