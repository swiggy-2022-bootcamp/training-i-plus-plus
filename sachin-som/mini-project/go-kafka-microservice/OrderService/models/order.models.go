package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	OrderID       primitive.ObjectID `json:"_id"            bson:"_id"`
	UserID        primitive.ObjectID `json:"user_id"        bson:"user_id"`
	OrderCart     Product            `json:"order_cart"     bson:"order_cart"`
	OrderedAt     time.Time          `json:"ordered_at"     bson:"ordered_at"`
	Bill          string             `json:"price"          bson:"price"`
	Discount      string             `json:"discount"       bson:"discount"`
	PaymentMethod string             `json:"payment_method" bson:"payment_method"`
	Status        string             `json:"status"         bson:"status"`
}
