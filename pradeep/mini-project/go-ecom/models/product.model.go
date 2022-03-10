package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Product_Name string             `json:"product_name" bson:"product_name"`
	Price        float64            `json:"price" bson:"price"`
	Rating       int                `json:"rating" bson:"rating"`
	Image        string             `json:"image" bson:"image"`
}
