package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Payment struct {
	PaymentID primitive.ObjectID `json:"_id"     bson:"_id"`
	Digital   bool               `json:"digital" bson:"digital"`
	COD       bool               `json:"cod"     bson:"cod"`
}
