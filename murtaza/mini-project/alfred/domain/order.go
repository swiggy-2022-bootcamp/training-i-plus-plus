package domain

import "alfred/utils/errs"

type Order struct {
	Id          int            `json:"order_id"`
	UserId      int            `json:"user_id"`
	OrderAmount float64        `json:"order_amount"`
	Items       map[string]int `json:"items"`
}

type OrderRepository interface {
	InsertOrder(Order) (*Order, *errs.AppError)
}

func NewOrder(userId int, orderAmount float64, items map[string]int) *Order {
	return &Order{
		UserId:      userId,
		OrderAmount: orderAmount,
		Items:       items,
	}
}
