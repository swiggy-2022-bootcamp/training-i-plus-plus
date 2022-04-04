package domain

import (
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/ayan/mini-project/shopping-app/order/utils/errs"
)

type OrderService interface {
	CreateOrder(Order) (*Order, *errs.AppError)
	FindById(string) (*Order, *errs.AppError)
}

type DefaultOrderService struct {
	OrderDB OrderRepositoryDB
}

func NewOrderService(orderDB OrderRepositoryDB) OrderService {
	return &DefaultOrderService{
		OrderDB: orderDB,
	}
}

func (osvc DefaultOrderService) CreateOrder(order Order) (*Order, *errs.AppError) {

	// calculate amount here
	order.Amount = 0

	u, err := osvc.OrderDB.Save(order)

	return u, err
}

func (osvc *DefaultOrderService) FindById(id string) (*Order, *errs.AppError) {

	order, err := osvc.OrderDB.FetchOrderById(id)
	return order, err
}
