package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        *string            `json:"name" validate:"required,min=2,max=100"`
	Price       *float64           `json:"price" validate:"required"`
	Stock_units *int               `json:stock_units`
	Product_id  string             `json:"product_id"`
	Seller_id   string
}
