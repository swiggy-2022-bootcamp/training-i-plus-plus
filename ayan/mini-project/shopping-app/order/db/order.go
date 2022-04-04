package db

import (
	"time"

	"order/domain"
)

type Order struct {
	Id        string             `bson:"id"`
	UserEmail string             `bson:"user_email"`
	ItemList  []domain.OrderItem `bson:"item_list"`
	Amount    int                `bson:"amount"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func NewOrder(userEmail string, itemList []domain.OrderItem, amount int) *Order {
	return &Order{
		UserEmail: userEmail,
		ItemList:  itemList,
		Amount:    amount,
	}
}
