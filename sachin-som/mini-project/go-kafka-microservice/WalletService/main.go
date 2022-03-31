package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/WalletService/controllers"
	"github.com/go-kafka-microservice/WalletService/database"
	"github.com/go-kafka-microservice/WalletService/services"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server *gin.Engine
	err    error
	ctx    context.Context

	mongoClient      *mongo.Client
	walletCollection *mongo.Collection
	walletServices   services.WalletServices
	walletController *controllers.WalletControllers
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

	// Create wallets Collection in MongoDB
	walletCollection = mongoClient.Database("WalletDB").Collection("wallets")

	// Initialize Services
	walletServices = services.NewWalletServiesImpl(walletCollection, ctx)

	// Initialize Controllers
	walletController = controllers.NewWalletControllers(walletServices)

	// server
	server = gin.Default()
}

func main() {
	defer mongoClient.Disconnect(ctx)

	base := server.Group("/v1/wallet")
	walletController.RegisterWalletRoutes(base)

	log.Fatal(server.Run(":8005"))
}
