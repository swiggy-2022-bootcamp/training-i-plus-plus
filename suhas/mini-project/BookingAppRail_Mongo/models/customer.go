package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
// const (
// 	// Admin holds the name of the admins collection
// 	CustomerCollection = "customer"
// )

// Article model
type Customer struct {
	Id             primitive.ObjectID`bson:"_id,omitempty"`
	Firstname      string              `bson:Fistname,omitempty`
	Lastname       string			   `bson:Lastname,omitempty`
	CustomerId     string              `bson:customerid,omitempty`
}