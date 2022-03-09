package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductUser struct {
	ProductID   primitive.ObjectID `json:"_id"          bson:"_id"`
	ProductName *string            `json:"product_name" bson:"product_name"`
	Price       int                `json:"price"        bson:"price"`
	Ratings     *uint              `json:"ratings"      bson:"ratings"`
	ImageUrl    *string            `json:"image_url"    bson:"image_url"`
}
