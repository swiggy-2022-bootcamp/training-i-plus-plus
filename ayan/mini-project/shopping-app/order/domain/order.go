package domain

import (
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/ayan/mini-project/shopping-app/order/utils/errs"
)

type Order struct {
	Id       string      `json:"name,omitempty"`
	ItemList []OrderItem `json:"item_list"`
	Amount   int         `json:"amount,omitempty"`
}

type OrderItem struct {
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func NewOrder(itemList []OrderItem, amount int) *Order {
	return &Order{
		ItemList: itemList,
		Amount:   amount,
	}
}

type OrderRepositoryDB interface {
	Save(Order) (*Order, *errs.AppError)
	FetchOrderById(string) (*Order, *errs.AppError)
}
