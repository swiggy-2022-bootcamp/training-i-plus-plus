package service

import (
	"Order-Service/config"
	errors "Order-Service/errors"
	"Order-Service/kafka"
	mockdata "Order-Service/model"
	"context"
	"encoding/json"
	"io"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var mongoURL string = config.MONGO_URL
var orderCollection *mongo.Collection
var i = 0

func init() {
	// Initialize a new mongo client with options
	client, _ = mongo.NewClient(options.Client().ApplyURI(mongoURL))
	orderCollection = client.Database("swiggy_mini").Collection("orders")
}

func PlaceOrder(body *io.ReadCloser) (result *mongo.InsertOneResult) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	var orderPlaced mockdata.Order
	json.NewDecoder(*body).Decode(&orderPlaced)

	orderPlaced.OrderDate = time.Now()
	orderPlaced.DeliveryDate = orderPlaced.OrderDate.AddDate(0, 0, 6)
	orderPlaced.Status = "confirmed"

	ctx, _ = context.WithTimeout(context.Background(), time.Second*10)
	result, _ = orderCollection.InsertOne(ctx, orderPlaced)

	ctx, _ = context.WithTimeout(context.Background(), time.Minute*10)
	kafka.Produce(ctx, []byte(strconv.Itoa(i)), []byte("order placed by user with id "+orderPlaced.UserId+" --- status: "+orderPlaced.Status))

	i++

	return
}

func GetOrders(userId string) (orders []mockdata.Order) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	//TODO: check if user exists
	ctx, _ = context.WithTimeout(context.Background(), time.Second*10)
	cursor, _ := orderCollection.Find(ctx, bson.M{"userid": userId})

	for cursor.Next(ctx) {
		var order mockdata.Order
		cursor.Decode(&order)
		orders = append(orders, order)
	}
	return
}

func OrderPayment(orderId string) (successMessage *string, err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return nil, errors.MalformedIdError()
	}

	result := orderCollection.FindOne(ctx, bson.M{"_id": objectId})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		return nil, errors.IdNotFoundError()
	}

	var order mockdata.Order
	result.Decode(&order)

	if order.Status == "payment done" || order.Status == "delivered" {
		return nil, errors.OrderAlreadyPaidForError()
	}

	order.Status = "payment done"

	_, error := orderCollection.UpdateByID(ctx, objectId, bson.M{"$set": order})

	if error != nil {
		return nil, errors.InternalServerError()
	}

	ctx, _ = context.WithTimeout(context.Background(), time.Minute*10)
	kafka.Produce(ctx, []byte(strconv.Itoa(i)), []byte("orderId: "+orderId+" --- status: "+order.Status))

	i++

	str := "order payment successful"
	successMessage = &str
	return
}

func DeliverOrder(orderId string) (successMessage *string, err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return nil, errors.MalformedIdError()
	}

	result := orderCollection.FindOne(ctx, bson.M{"_id": objectId})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		return nil, errors.IdNotFoundError()
	}

	var order mockdata.Order
	result.Decode(&order)

	if order.Status == "confirmed" {
		return nil, errors.PaymentIncompleteError()
	}

	if order.Status == "delivered" {
		return nil, errors.OrderAlreadyDeliveredError()
	}

	order.Status = "delivered"

	_, error := orderCollection.UpdateByID(ctx, objectId, bson.M{"$set": order})
	if error != nil {
		return nil, errors.InternalServerError()
	}

	ctx, _ = context.WithTimeout(context.Background(), time.Minute*10)
	kafka.Produce(ctx, []byte(strconv.Itoa(i)), []byte("orderId: "+orderId+" --- status: "+order.Status))

	i++

	str := "order delivered"
	successMessage = &str
	return
}
