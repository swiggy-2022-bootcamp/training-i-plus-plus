package db

import (
	"order.akash.com/model"
)

type OrderRepository interface {
	Connect()
	FindAll() []model.Order
	SaveOrder(model.Order)
}
