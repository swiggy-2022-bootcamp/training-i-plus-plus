package domain

import (
	"order/utils/errs"
)

type Order struct {
	Id        string      `json:"id,omitempty"`
	UserEmail string      `json:"user_email"`
	ItemList  []OrderItem `json:"item_list"`
	Amount    int         `json:"amount,omitempty"`
}

type OrderItem struct {
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func NewOrder(userEmail string, itemList []OrderItem, amount int) *Order {
	return &Order{
		UserEmail: userEmail,
		ItemList:  itemList,
		Amount:    amount,
	}
}

type OrderRepositoryDB interface {
	Save(Order) (*Order, *errs.AppError)
	FetchOrderById(string) (*Order, *errs.AppError)
}
