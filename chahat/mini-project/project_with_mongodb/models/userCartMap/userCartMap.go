 //package usermodel

// type User struct {
// 	ID string `json:"id"`
//     Email string `json:"email"`
// 	FirstName string `json:"firstName"`
// 	LastName string `json:"lastName"`
// 	Password string `json:"password"`
// }
//package models

// import (
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type Address struct {
// 	//TODO: add it to model and use "bson inline" during unmarshal
// 	Street  string
// 	City    string
// 	Country string
// }

// type User struct {
// 	Id         primitive.ObjectID `bson:"_id"`
// 	FirstName  string             `json:"firstname" validate:"min=2,max=100,required"`
// 	LastName   string             `json:"lastname" validate:"min=2,max=100,required"`
// 	Password   string             `json:"password" validate:"min=6,required"`
// 	Email      string             `json:"email" validate:"email,required"`
// 	Phone      string             `json:"phone" validate:"required"`
// //	Token      string             `json:"token"`
// 	Created_at time.Time          `json:"created_at"`
// 	User_id    string             `json:"user_id"`
// }
package usermodel

import(
	"time"
"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	ID				primitive.ObjectID		`bson:"_id"`
	Cart_id		    string					   `json:"cart_id"`
	User_id			string					`json:"user_id"`
}