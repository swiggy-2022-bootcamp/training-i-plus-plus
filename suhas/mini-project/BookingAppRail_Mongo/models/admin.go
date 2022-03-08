package models

import (
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Article model
type Admin struct {
	Id         primitive.ObjectID  `bson:"_id,omitempty"`
	Name       string              `bson:"name,omitempty"`
	AdminId    string              `bson:"adminid,omitempty"`
}
