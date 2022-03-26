package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Name        string               `json:"name,omitempty" validate:"required"`
	Email       string               `json:"email,omitempty" validate:"required"`
	PurchasedID []primitive.ObjectID `json:"purchasedticket,omitempty"`
}
