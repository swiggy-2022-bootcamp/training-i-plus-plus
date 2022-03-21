package controller

import (
	"context"
	"encoding/json"
	"net/http"
	mockdata "src/model"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

	connection := client.Database("swiggy_mini").Collection("orders")
	ctx, _ = context.WithTimeout(context.Background(), time.Second*10)
	result, _ := connection.InsertOne(ctx, orderPlaced)

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
