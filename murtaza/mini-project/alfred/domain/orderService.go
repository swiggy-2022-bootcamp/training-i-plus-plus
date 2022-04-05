package domain

import (
	"alfred/utils/errs"
)

type OrderService interface {
	CreateOrder(userId int, orderAmount float64, items map[string]int) (*Order, *errs.AppError)
}

type orderService struct {
	orderRepository OrderRepository
	producer        OrderProducer
}

func (os orderService) CreateOrder(userId int, orderAmount float64, items map[string]int) (*Order, *errs.AppError) {
	newOrder := NewOrder(userId, orderAmount, items)
	order, err := os.orderRepository.InsertOrder(*newOrder)
	if err != nil {
		return nil, err
	}
	os.producer.SendOrderAmount(order.Id, userId, orderAmount)
	return order, nil
}

func NewOrderService(repository OrderRepository, producer OrderProducer) OrderService {
	return &orderService{
		orderRepository: repository,
		producer:        producer,
	}
}
