package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id"           bson:"_id"`
	Fullname string             `json:"fullname"      bson:"fullname"      validate:"required"`
	Email    string             `json:"email"         bson:"email"         validate:"required"`
	Phone    string             `json:"phone"         bson:"phone"         validate:"required"`
	Password string             `json:"password"      bson:"password"      validate:"required"`
}
