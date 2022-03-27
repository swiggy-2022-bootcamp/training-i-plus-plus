package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	FirstName     *string            `json:"first_name" validate:"required,min=1,max=20"`
	LastName      *string            `json:"last_name" validate:"required,min=1,max=20"`
	Password      *string            `json:"Password" validate:"required,min=8,max=40"`
	Email         *string            `json:"email" validate:"email,required"`
	Phone         *string            `json:"phone" validate:"required"`
	Token         *string            `json:"token"`
	Refresh_Token *string            `json:"refresh_token"`
	User_id       string             `json:"user_id"`
}
