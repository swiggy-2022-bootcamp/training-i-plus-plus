package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID             primitive.ObjectID `bson:"_id"`
	Items          []interface{}
	Payment_method string    `json:"payment_method" validate:"eq=ONLINE|eq=COD"`
	Payment_status string    `json:"payment_status" validate:"eq=PAID|eq=UNPAID"`
	Created_at     time.Time `json:"created_at"`
	Order_id       string    `json:"invoice_id"`
	User_id        string    `json:"user_id"`
}
