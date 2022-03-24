package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/InventoryService/controllers"
	"github.com/go-kafka-microservice/InventoryService/database"
	"github.com/go-kafka-microservice/InventoryService/services"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	err         error
	ctx         context.Context
	server      *gin.Engine
	mongoClient *mongo.Client

	inventoryController *controllers.InventoryControllers
	inventoryService    services.InventoryServices
	inventoryCollection *mongo.Collection
	productCollection   *mongo.Collection
)

func init() {
	// Load Dotenv file
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error Loading in .env file: ", err.Error())
	}

	// Create Context
	ctx = context.TODO()

	// MongoClient Connection
	mongoClient = database.SetUpDatabase(ctx)

	// Create User Collection in MongoDB
	inventoryCollection = mongoClient.Database("InventoryDB").Collection("inventory")
	productCollection = mongoClient.Database("InventoryDB").Collection("product")

	// Initialize layers
	inventoryService = services.NewInventoryService(inventoryCollection, productCollection, ctx)
	inventoryController = controllers.NewInventoryControllers(inventoryService)

	// Initialize gin server
	server = gin.Default()
}

func main() {
	defer mongoClient.Disconnect(ctx)

	// Define Base path and register routes
	basePath := server.Group("/v1/inventory")
	inventoryController.RegisterInventoryRoutes(basePath)

	// start server
	log.Fatal(server.Run(":8002"))
}
