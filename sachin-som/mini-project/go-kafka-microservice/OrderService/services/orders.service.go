package services

import (
	"github.com/go-kafka-microservice/OrderService/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderServices interface {
	SaveOrder(models.Order) error
	GetOrders(primitive.ObjectID) ([]*models.Order, error)
}
