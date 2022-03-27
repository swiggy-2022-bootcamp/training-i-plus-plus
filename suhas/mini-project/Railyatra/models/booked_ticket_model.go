package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Passengers struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type BookedTicket struct {
	Id              primitive.ObjectID `json:"id,omitempty"`
	Train_id        primitive.ObjectID `json:"trainid,omitempty"`
	Departure       string             `json:"departure"`
	Arrival         string             `json:"arrival"`
	Departure_time  string             `json:"departure_time"`
	Arrival_time    string             `json:"arrival_time"`
	Amount_paid     int                `json:"amount_paid"`
	User_id         primitive.ObjectID `json:"userid,omitempty"`
	Passengers_info []Passengers       `json:"passenger"`
}
