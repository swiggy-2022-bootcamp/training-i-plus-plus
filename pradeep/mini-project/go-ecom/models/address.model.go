package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Address_Line1 string             `json:"address_line1" bson:"address_line1"`
	Address_Line2 string             `json:"address_line2" bson:"address_line2"`
	City          string             `json:"city" bson:"city"`
	State         string             `json:"state" bson:"state"`
	Country       string             `json:"country" bson:"country"`
	Zip_Code      string             `json:"zip_code" bson:"zip_code"`
	User_Id       string             `json:"user_id" bson:"user_id"`
}
