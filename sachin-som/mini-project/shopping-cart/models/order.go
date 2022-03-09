package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	OrderID       primitive.ObjectID `json:"_id"            bson:"_id"`
	OrderCart     []ProductUser      `json:"order_cart"     bson:"order_cart"`
	OrderedAt     time.Time          `json:"ordered_at"     bson:"ordered_at"`
	Price         int                `json:"price"          bson:"price"`
	Discount      int                `json:"discount"       bson:"discount"`
	PaymentMethod Payment            `json:"payment_method" bson:"payment_method"`
}
