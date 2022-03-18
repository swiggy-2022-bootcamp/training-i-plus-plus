package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Passengers struct {
	Name string
	Age  int
}

type BookedTicket struct {
	Id              primitive.ObjectID `json:"adminId,omitempty"`
	Train_id        primitive.ObjectID `json:"trainId,omitempty"`
	Departure       string             `json:"departure"`
	Arrival         string             `json:"arrival"`
	Departure_time  time.Time          `json:"departure_time"`
	Arrival_time    time.Time          `json:"arrival_time"`
	Amount_paid     int                `json:"amount_paid"`
	User_id         primitive.ObjectID `json:"user_id,omitempty"`
	Passengers_info []Passengers       `json:"passenger"`
}
