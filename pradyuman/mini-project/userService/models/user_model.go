package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    UserId       	primitive.ObjectID `json:"userid,omitempty"`
    Name     	string             `json:"name,omitempty" validate:"required"`
	Email	 	string			   `json:"email" validate:"email,required"`
    Password 	string			   `json:"password" validate:"required,min=8"`
	Phone	 	string			   `json:"phone" validate:"required"`
    Role		string			   `json:"role" validate:"required"`
}