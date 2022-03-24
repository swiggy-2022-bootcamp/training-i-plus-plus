package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/UserService/controllers"
	"github.com/go-kafka-microservice/UserService/database"
	"github.com/go-kafka-microservice/UserService/services"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	err         error
	ctx         context.Context
	server      *gin.Engine
	mongoClient *mongo.Client

	userController *controllers.UserControllers
	userService    services.UserService
	userCollection *mongo.Collection
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
	userCollection = mongoClient.Database("UserDB").Collection("users")

	// Initialize layers
	userService = services.NewUserServiceImpl(userCollection, ctx)
	userController = controllers.NewUserControllers(userService)

	// Initialize gin server
	server = gin.Default()
}

func main() {
	defer mongoClient.Disconnect(ctx)

	// Define Base path and register routes
	basePath := server.Group("/v1/users")
	userController.RegisterUserRoutes(basePath)

	// start server
	log.Fatal(server.Run(":8001"))
}
