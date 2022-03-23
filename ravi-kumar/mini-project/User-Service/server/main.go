package main

import (
	"User-Service/config"
	"User-Service/controller"
	"context"
	"log"
	"net/http"
	"net/url"
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

	//go kafka.Consume(ctx)

	router := mux.NewRouter()
	router.HandleFunc("/users", controller.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", controller.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{userId}", controller.GetUserById).Methods(http.MethodGet)
	router.HandleFunc("/users/{userId}", controller.UpdateUserById).Methods(http.MethodPut)
	router.HandleFunc("/users/{userId}", controller.DeleteUserbyId).Methods(http.MethodDelete)

	log.Print("User Service: Starting server at port ", config.USER_SERVICE_SERVER_PORT)
	http.ListenAndServe(":5004", router)
}

func GetClient() *mongo.Client {
	return client
}
