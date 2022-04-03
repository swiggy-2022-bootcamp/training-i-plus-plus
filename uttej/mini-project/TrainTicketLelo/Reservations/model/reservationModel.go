package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reservation struct {
	Id            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId        string             `json: "userId" bson: "userId"`
	TrainIDs      []string           `json: "ticketids" bson: "ticketids"`
	Amount        float64            `json: "amount" bson: "amount"`
	PurchaseDate  time.Time          `json: "purchaseDate" bson: "purchaseDate"`
	DepartureDate time.Time          `json: "departureDate" bson: "departureDate"`
	Status        string             `json:"status" bson: "status"`
}
