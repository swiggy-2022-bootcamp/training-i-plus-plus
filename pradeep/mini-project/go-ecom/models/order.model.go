package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Order_Cart     []ProductUser      `json:"order_cart" bson:"order_cart"`
	Ordered_At     time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price          float64            `json:"price" bson:"price"`
	Discount       float64            `json:"discount" bson:"discount"`
	Payment_Method Payment            `json:"payment_method" bson:"payment_method"`
}
