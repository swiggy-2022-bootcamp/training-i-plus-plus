package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	UserName string             `json:"username,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Role     string             `json:"role,omitempty" validate:"required"`
}

type Authentication struct {
	UserName string
	Password string
}
