package domain

import "alfred/utils/errs"

type OrderService interface {
	CreateOrder(userId int, orderAmount float64, items map[string]int) (*Order, *errs.AppError)
}

type orderService struct {
	orderRepository OrderRepository
}

func (os orderService) CreateOrder(userId int, orderAmount float64, items map[string]int) (*Order, *errs.AppError) {
	newOrder := NewOrder(userId, orderAmount, items)
	return os.orderRepository.InsertOrder(*newOrder)
}

func NewOrderService(repository OrderRepository) OrderService {
	return &orderService{
		orderRepository: repository,
	}
}
