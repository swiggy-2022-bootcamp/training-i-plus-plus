package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role int

const (
	Admin     Role = 1
	Traveller Role = 2
)

func IsValidRole(role int) bool {
	return (role == 1 || role == 2)
}

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Fullname string             `json: "fullName" bson: "fullName"`
	UserName string             `json: "userName" bson: "userName"`
	Password string             `json: "password" bson: "password"`
	Address  string             `json: "address" bson: "address"`
	Role     Role               `json:"role" bson: "role"`
}

type Login struct {
	UserName string `json: "userName" bson: "userName"`
	Password string `json: "password" bson: "password"`
}
