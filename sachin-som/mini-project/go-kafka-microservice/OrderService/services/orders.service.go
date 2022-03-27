package services

import "github.com/go-kafka-microservice/OrderService/models"

type OrderServices interface {
	SaveOrder(models.Order) error
}
