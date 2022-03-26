package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Passengers struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Purchased struct {
	Train_id        primitive.ObjectID `json:"trainid,omitempty"`
	Departure       string             `json:"departure"`
	Arrival         string             `json:"arrival"`
	Departure_time  string             `json:"departure_time"`
	Arrival_time    string             `json:"arrival_time"`
	Cost            int                `json:"cost"`
	User_id         primitive.ObjectID `json:"userid,omitempty"`
	Passengers_info []Passengers       `json:"passenger"`
}
