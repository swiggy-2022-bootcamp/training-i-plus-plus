package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	mongoId     primitive.ObjectID `bson:"_id,omitempty"`
	Id          string             `bson:"id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Quantity    int                `bson:"quantity"`
	Price       int                `bson:"price"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

func NewProduct(name string, description string, quantity int, price int) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Quantity:    quantity,
		Price:       price,
	}
}
