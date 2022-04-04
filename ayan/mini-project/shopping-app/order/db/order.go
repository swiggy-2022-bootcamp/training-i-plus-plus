package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	mongoId   primitive.ObjectID `bson:"_id,omitempty"`
	Id        string             `bson:"id"`
	ItemList  []OrderItem        `bson:"item_list"`
	Amount    int                `bson:"amount"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type OrderItem struct {
	ProductId string `bson:"product_id"`
	Quantity  int    `bson:"quantity"`
}

func NewOrder(itemList []OrderItem, amount int) *Order {
	return &Order{
		ItemList: itemList,
		Amount:   amount,
	}
}
