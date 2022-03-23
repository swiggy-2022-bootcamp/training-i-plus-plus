package main

import (
	"Inventory-Service/config"
	"Inventory-Service/controller"
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

	router.HandleFunc("/catalog", controller.CreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/catalog", controller.GetCatalog).Methods(http.MethodGet)
	router.HandleFunc("/catalog/{productId}", controller.GetProductById).Methods(http.MethodGet)
	router.HandleFunc("/catalog/{productId}", controller.UpdateProductById).Methods(http.MethodPut)
	router.HandleFunc("/catalog/{productId}", controller.DeleteProductbyId).Methods(http.MethodDelete)

	log.Print("Inventory Service: Starting server at port ", config.INVENTORY_SERVICE_SERVER_PORT)
	http.ListenAndServe(":5002", router)
}
