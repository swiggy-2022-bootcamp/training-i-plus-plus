package domain

import "alfred/utils/errs"

type Cart struct {
	Id     int            `json:"id"`
	UserId int            `json:"user_id"`
	Items  map[string]int `json:"items"`
}

type CartRepository interface {
	AddToCart(int, map[string]int) *errs.AppError
	GetCart(int) (*Cart, *errs.AppError)
	RemoveCart(int) *errs.AppError
}

func NewCart(userId int, items map[string]int) *Cart {
	return &Cart{
		UserId: userId,
		Items:  items,
	}
}
