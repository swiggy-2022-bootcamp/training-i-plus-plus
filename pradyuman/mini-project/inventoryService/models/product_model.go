package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct{
	ProductId   primitive.ObjectID `json:"productid,omitempty"`
    Name        string         `json:"name" validate:"required"`
	SellerId	string		   `json:"sellerid" validate:"required"`
	Price		int			   `json:"price" validate:"required"`
	Quantity    int            `json:"quantity" validate:"required"`
}