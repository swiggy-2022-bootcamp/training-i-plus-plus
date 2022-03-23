package controller

import (
	"Order-Service/config"
	mockdata "Order-Service/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"Order-Service/kafka"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var mongoURL string = config.MONGO_URL

var i int = 0

func init() {
	// Initialize a new mongo client with options
	client, _ = mongo.NewClient(options.Client().ApplyURI(mongoURL))
}

func PlaceOrder(res http.ResponseWriter, req *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	var orderPlaced mockdata.Order
	json.NewDecoder(req.Body).Decode(&orderPlaced)

	orderPlaced.OrderDate = time.Now()
	orderPlaced.DeliveryDate = orderPlaced.OrderDate.AddDate(0, 0, 6)
	orderPlaced.Status = "confirmed"

	connection := client.Database("swiggy_mini").Collection("orders")
	ctx, _ = context.WithTimeout(context.Background(), time.Second*10)
	result, _ := connection.InsertOne(ctx, orderPlaced)

	ctx, _ = context.WithTimeout(context.Background(), time.Minute*10)
	kafka.Produce(ctx, []byte(strconv.Itoa(i)), []byte("order placed by user with id "+orderPlaced.UserId+" --- status: "+orderPlaced.Status))

	i++

	res.Header().Add("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}

func GetOrders(res http.ResponseWriter, req *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	vars := mux.Vars(req)
	userId := vars["userId"]
	connection := client.Database("swiggy_mini").Collection("orders")
	ctx, _ = context.WithTimeout(context.Background(), time.Second*10)
	cursor, _ := connection.Find(ctx, bson.M{"userid": userId})

	var orders []mockdata.Order

	for cursor.Next(ctx) {
		var order mockdata.Order
		cursor.Decode(&order)
		orders = append(orders, order)
	}

	res.Header().Add("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(orders)
}

func OrderPayment(res http.ResponseWriter, req *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	vars := mux.Vars(req)
	orderId := vars["orderId"]

	connection := client.Database("swiggy_mini").Collection("orders")
	ctx, _ = context.WithTimeout(context.Background(), time.Second*10)
	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode("Malformed order id")
		return
	}

	result := connection.FindOne(ctx, bson.M{"_id": objectId})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode("order with given id not found")
		return
	}
	var order mockdata.Order
	result.Decode(&order)

	if order.Status == "payment done" || order.Status == "delivered" {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode("Order has already been paid for. Aborting this Payment.")
		return
	}

	order.Status = "payment done"

	_, error := connection.UpdateByID(ctx, objectId, bson.M{"$set": order})
	if error != nil {
		fmt.Println(error)
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode("Internal server error")
		return
	}

	ctx, _ = context.WithTimeout(context.Background(), time.Minute*10)
	kafka.Produce(ctx, []byte(strconv.Itoa(i)), []byte("orderId: "+orderId+" --- status: "+order.Status))
	i++

	res.Header().Add("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode("order payment successful")
}

func DeliverOrder(res http.ResponseWriter, req *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	vars := mux.Vars(req)
	orderId := vars["orderId"]

	connection := client.Database("swiggy_mini").Collection("orders")
	ctx, _ = context.WithTimeout(context.Background(), time.Second*10)
	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode("Malformed order id")
		return
	}

	result := connection.FindOne(ctx, bson.M{"_id": objectId})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode("order with given id not found")
		return
	}
	var order mockdata.Order
	result.Decode(&order)

	if order.Status == "confirmed" {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode("Order payment not done. Current delivery aborted.")
		return
	}

	if order.Status == "delivered" {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode("Order has already been delivered. Current delivery aborted.")
		return
	}

	order.Status = "delivered"

	_, error := connection.UpdateByID(ctx, objectId, bson.M{"$set": order})
	if error != nil {
		fmt.Println(error)
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode("Internal server error")
		return
	}

	ctx, _ = context.WithTimeout(context.Background(), time.Minute*10)
	kafka.Produce(ctx, []byte(strconv.Itoa(i)), []byte("orderId: "+orderId+" --- status: "+order.Status))
	i++

	res.Header().Add("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode("order delivered")
}

//confirmed -> payment done -> delivered
