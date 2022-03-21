package main

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"src/config"
	"src/controller"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var mongoURL = "mongodb+srv://rshantharaju:" + url.QueryEscape("Ravi@1999") + "@cluster0.05bio.mongodb.net/test"

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// Initialize a new mongo client with options
	client, _ := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	_ = client.Connect(ctx)

	router := mux.NewRouter()
	router.HandleFunc("/users", controller.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", controller.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{userId}", controller.GetUserById).Methods(http.MethodGet)
	router.HandleFunc("/users/{userId}", controller.UpdateUserById).Methods(http.MethodPut)
	router.HandleFunc("/users/{userId}", controller.DeleteUserbyId).Methods(http.MethodDelete)

	router.HandleFunc("/catalog", controller.CreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/catalog", controller.GetCatalog).Methods(http.MethodGet)
	router.HandleFunc("/catalog/{productId}", controller.GetProductById).Methods(http.MethodGet)
	router.HandleFunc("/catalog/{productId}", controller.UpdateProductById).Methods(http.MethodPut)
	router.HandleFunc("/catalog/{productId}", controller.DeleteProductbyId).Methods(http.MethodDelete)

	router.HandleFunc("/order", controller.PlaceOrder).Methods(http.MethodPost)
	router.HandleFunc("/{userId}/order", controller.GetOrders).Methods(http.MethodGet)

	log.Print("Starting server at port ", config.SERVER_PORT)
	http.ListenAndServe(":5001", router)
}

func GetClient() *mongo.Client {
	return client
}
