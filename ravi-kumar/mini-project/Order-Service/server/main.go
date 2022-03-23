package main

import (
	"Order-Service/config"
	"Order-Service/controller"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// Initialize a new mongo client with options
	client, _ := mongo.NewClient(options.Client().ApplyURI(config.MONGO_URL))
	_ = client.Connect(ctx)

	//go kafka.Consume(ctx)

	router := mux.NewRouter()

	router.HandleFunc("/order", controller.PlaceOrder).Methods(http.MethodPost)
	router.HandleFunc("/{userId}/order", controller.GetOrders).Methods(http.MethodGet)
	router.HandleFunc("/order/{orderId}/payment", controller.OrderPayment).Methods(http.MethodPost)
	router.HandleFunc("/order/{orderId}/deliver", controller.DeliverOrder).Methods(http.MethodPost)

	log.Print("Order Service: Starting server at port ", config.ORDER_SERVICE_SERVER_PORT)
	http.ListenAndServe(":5003", router)
}
