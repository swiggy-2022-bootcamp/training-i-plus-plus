package mockdata

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId       string             `json: "userId" bson: "userId"`
	Items        []string           `json: "items" bson: "items"`
	Amount       float64            `json: "amount" bson: "amount"`
	OrderDate    time.Time          `json: "orderDate" bson: "orderDate"`
	DeliveryDate time.Time          `json: "deliveryDate" bson: "deliveryDate"`
	Status       string             `json:"status" bson: "status"`
}
