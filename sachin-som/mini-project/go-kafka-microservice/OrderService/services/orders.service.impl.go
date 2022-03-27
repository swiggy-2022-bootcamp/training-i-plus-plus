package services

import (
	"context"

	"github.com/go-kafka-microservice/OrderService/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderServicesImpl struct {
	OrderCollection *mongo.Collection
	Ctx             context.Context
}

// Function to create new instance of OrderServiceImpl
func NewOrderServiceImpl(orderCollection *mongo.Collection, ctx context.Context) *OrderServicesImpl {
	return &OrderServicesImpl{
		OrderCollection: orderCollection,
		Ctx:             ctx,
	}
}

// Function to save orders to OrderCollection
func (os *OrderServicesImpl) SaveOrder(order *models.Order) error {
	return nil
}
