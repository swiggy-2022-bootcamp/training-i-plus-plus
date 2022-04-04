package domain

import "github.com/swiggy-2022-bootcamp/training-i-plus-plus/murtaza/mini-project/inventory/utils/errs"

type Item struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func NewItem(name, description string, quantity int) *Item {
	return &Item{
		Name:        name,
		Description: description,
		Quantity:    quantity,
	}
}

type ItemRepository interface {
	InsertItem(Item) (Item, *errs.AppError)
	FindItemById(int) (*Item, *errs.AppError)
	FindItemByName(string) (*Item, *errs.AppError)
	DeleteItemById(int) *errs.AppError
	UpdateItem(Item) (*Item, *errs.AppError)
}
