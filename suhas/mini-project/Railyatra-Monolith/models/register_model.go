package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Register struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Username string             `json:"username,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Group    string             `json:"group,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
}
