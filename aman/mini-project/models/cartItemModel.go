package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartItem struct {
	ID         primitive.ObjectID `bson:"_id"`
	Product_id string             `json:"product_id"`
	Quantity   *int               `json:"quantity" validate:"required"`
	User_id    string             `json:"user_id"`
}
