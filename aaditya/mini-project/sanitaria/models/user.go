package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct{
	Id			primitive.ObjectID	`json:"id,omitempty"`
	Name 		string	`json:"name"`
	EmailId 	string  `json:"emailId"`
	Age 		int		`json:"age"`
	Address		string  `json:"address"`
}