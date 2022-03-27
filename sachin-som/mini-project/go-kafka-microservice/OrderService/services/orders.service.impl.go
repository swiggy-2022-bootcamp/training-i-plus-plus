package services

import (
	"context"
	"errors"

	"github.com/go-kafka-microservice/OrderService/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func (os *OrderServicesImpl) GetOrders(usrID primitive.ObjectID) ([]*models.Order, error) {
	var orders []*models.Order
	filter := bson.D{bson.E{Key: "user_id", Value: usrID}}
	cursor, err := os.OrderCollection.Find(os.Ctx, filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(os.Ctx) {
		var _order models.Order
		err := cursor.Decode(&_order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &_order)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(os.Ctx)
	if len(orders) == 0 {
		return nil, errors.New("Orders not fuond.")
	}
	return orders, nil
}
