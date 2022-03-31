package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	pb "github.com/go-kafka-microservice/AuthProto"
	"github.com/go-kafka-microservice/InventoryService/controllers"
	"github.com/go-kafka-microservice/InventoryService/database"
	"github.com/go-kafka-microservice/InventoryService/middleware"
	"github.com/go-kafka-microservice/InventoryService/services"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	err         error
	ctx         context.Context
	server      *gin.Engine
	mongoClient *mongo.Client

	grpcConn            *grpc.ClientConn
	inventoryCollection *mongo.Collection
	productCollection   *mongo.Collection
	authProtoClient     pb.AuthServicesClient
	inventoryService    services.InventoryServices
	InventoryMiddleware *middleware.InventoryMiddleware
	inventoryController *controllers.InventoryControllers
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

	// Create gRPC Client
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	grpcConn, err = grpc.Dial("localhost:8000", opts...)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize AuthProtoClient
	authProtoClient = pb.NewAuthServicesClient(grpcConn)

	// Create User Collection in MongoDB
	inventoryCollection = mongoClient.Database("InventoryDB").Collection("inventory")
	productCollection = mongoClient.Database("InventoryDB").Collection("product")

	// Initialize layers
	inventoryService = services.NewInventoryService(inventoryCollection, productCollection, authProtoClient, ctx)
	InventoryMiddleware = middleware.NewInventoryMiddleware(inventoryService)
	inventoryController = controllers.NewInventoryControllers(inventoryService)

	// Initialize gin server
	server = gin.Default()

}

func main() {
	defer mongoClient.Disconnect(ctx)

	// Register Middlwares
	server.Use(InventoryMiddleware.AuthorizeUser())

	// Define Base path and register routes
	basePath := server.Group("/v1/inventory")
	inventoryController.RegisterInventoryRoutes(basePath)

	// start server
	log.Fatal(server.Run(":8002"))
}
