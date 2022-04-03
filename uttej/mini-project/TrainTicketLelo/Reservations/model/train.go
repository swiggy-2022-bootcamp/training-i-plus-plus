package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Train struct {
	Id               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TrainNumber      string             `json: "trainNumber" bson: "trainNumber"`
	TrainName        string             `json: "trainName" bson: "trainNname"`
	Price            int                `json: "price" bson: "price"`
	Source           string             `json: "source" bson: "source"`
	Destination      string             `json: "destination" bson: "destination"`
	AvailableTickets int                `json: "availableTickets" bson: "availableTickets"`
}
