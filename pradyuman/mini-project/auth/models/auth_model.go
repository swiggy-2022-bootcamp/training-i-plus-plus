package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Login struct {
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

type User struct {
    UserId      primitive.ObjectID `json:"userid,omitempty"`
    Name     	string             `json:"name,omitempty" validate:"required"`
	Email	 	string			   `json:"email" validate:"email,required"`
    Password 	string			   `json:"password" validate:"required,min=8"`
	Role		string			   `json:"role" validate:"required"`
	Phone	 	string			   `json:"phone" validate:"required"`
}