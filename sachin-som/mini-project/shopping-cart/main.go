package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sachinsom93/shopping-cart/controllers"
	"github.com/sachinsom93/shopping-cart/database"
	"github.com/sachinsom93/shopping-cart/services"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server         *gin.Engine
	userService    services.UserServices
	userController controllers.UserController
	ctx            context.Context
	userCollection *mongo.Collection
	mongoClient    *mongo.Client
	err            error
)

func init() {
	// Load .dotenv
	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error Loading in .env file: ", err.Error())
	}

	// Context to handle timeouts
	ctx = context.TODO()

	// Connnect to mongoDB
	mongoClient = database.SetUpDB(ctx)

	// Create Collections
	userCollection = mongoClient.Database("shopping-cart").Collection("users")

	// initiaze layers
	userService = services.NewUserService(userCollection, ctx)
	userController = controllers.NewUserController(userService)
	server = gin.Default()
}

func main() {
	defer mongoClient.Disconnect(ctx)

	basePath := server.Group("/v1")
	userController.RegisterUserRoutes(basePath)

	log.Fatal(server.Run(":9090"))
}
