package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	invcontroller "github.com/sachinsom93/shopping-cart/controllers/inventory"
	controllers "github.com/sachinsom93/shopping-cart/controllers/user"
	"github.com/sachinsom93/shopping-cart/database"
	"github.com/sachinsom93/shopping-cart/middleware"
	invservices "github.com/sachinsom93/shopping-cart/services/inventory"
	services "github.com/sachinsom93/shopping-cart/services/user"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	err         error
	ctx         context.Context
	server      *gin.Engine
	mongoClient *mongo.Client

	userService    services.UserServices
	userController controllers.UserController
	userCollection *mongo.Collection

	inventoryService    invservices.InventoryServices
	inventoryController invcontroller.InventoryController
	inventoryCollection *mongo.Collection
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
	inventoryCollection = mongoClient.Database("shopping-cart").Collection("inventories")

	// initiaze layers
	userService = services.NewUserService(userCollection, ctx)
	userController = controllers.NewUserController(userService)
	inventoryService = invservices.NewInventoryService(inventoryCollection, ctx)
	inventoryController = *invcontroller.NewInventoryController(inventoryService)

	// Don't need color console for logging to a file
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// Initialize server
	server = gin.Default()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// gin.DefaultWriter = file writer and os.Stdout
	server.Use(gin.LoggerWithFormatter(middleware.FormatLogger))

}

func main() {
	defer mongoClient.Disconnect(ctx)

	basePath := server.Group("/v1")
	userController.RegisterUserRoutes(basePath)
	inventoryController.RegisterInventoryRoutes(basePath)

	log.Fatal(server.Run(":9090"))
}
