package repository

import (
	"Order-Service/config"
	"Order-Service/errors"
	mockdata "Order-Service/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var mongoURL string = config.MONGO_URL
var orderCollection *mongo.Collection

func init() {
	// Initialize a new mongo client with options
	client, _ = mongo.NewClient(options.Client().ApplyURI(mongoURL))
	orderCollection = client.Database("swiggy_mini").Collection("orders")
}

type IMongoDAO interface {
	MongoPlaceOrder(orderPlaced mockdata.Order) string
	MongoGetOrderByUserId(userId string) (orders []mockdata.Order, err error)
	MongoGetOrderByOrderId(orderId primitive.ObjectID) (order *mockdata.Order, err error)
	MongoUpdateOrderByOrderId(orderId primitive.ObjectID, order mockdata.Order) (updatedOrder *mockdata.Order, err error)
	MongoDeliverOrderByOrderId(orderId primitive.ObjectID, order mockdata.Order) (updatedOrder *mockdata.Order, err error)
	MongoDeleteOrderById(orderId primitive.ObjectID)
}

type MongoDAO struct {
}

func (dao *MongoDAO) MongoDeleteOrderById(orderId primitive.ObjectID) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	orderCollection.DeleteOne(ctx, bson.M{"_id": orderId})
	return
}

func (dao *MongoDAO) MongoPlaceOrder(orderPlaced mockdata.Order) string {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	result, _ := orderCollection.InsertOne(ctx, orderPlaced)

	orderId := result.InsertedID.(primitive.ObjectID).Hex()
	return orderId
}

func (dao *MongoDAO) MongoGetOrderByUserId(userId string) (orders []mockdata.Order, err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	cursor, _ := orderCollection.Find(ctx, bson.M{"userid": userId})

	for cursor.Next(ctx) {
		var order mockdata.Order
		cursor.Decode(&order)
		orders = append(orders, order)
	}

	return
}

func (dao *MongoDAO) MongoGetOrderByOrderId(orderId primitive.ObjectID) (order *mockdata.Order, err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	result := orderCollection.FindOne(ctx, bson.M{"_id": orderId})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		return nil, errors.IdNotFoundError()
	}
	result.Decode(&order)
	return
}

func (dao *MongoDAO) MongoUpdateOrderByOrderId(orderId primitive.ObjectID, order mockdata.Order) (updatedOrder *mockdata.Order, err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	_, err = orderCollection.UpdateByID(ctx, orderId, bson.M{"$set": order})
	return
}

func (dao *MongoDAO) MongoDeliverOrderByOrderId(orderId primitive.ObjectID, order mockdata.Order) (updatedOrder *mockdata.Order, err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	_, err = orderCollection.UpdateByID(ctx, orderId, bson.M{"$set": order})
	return
}
