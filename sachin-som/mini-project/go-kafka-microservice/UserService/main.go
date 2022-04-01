package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	pb "github.com/go-kafka-microservice/AuthProto"
	"github.com/go-kafka-microservice/UserService/controllers"
	"github.com/go-kafka-microservice/UserService/database"
	_ "github.com/go-kafka-microservice/UserService/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/go-kafka-microservice/UserService/services"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	err         error
	ctx         context.Context
	server      *gin.Engine
	mongoClient *mongo.Client
	grpcConn    *grpc.ClientConn

	userController    *controllers.UserControllers
	authServiceClient pb.AuthServicesClient
	userService       services.UserService
	userCollection    *mongo.Collection
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

	// Create gRPC Client
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	grpcConn, err = grpc.Dial("localhost:8000", opts...)
	if err != nil {
		log.Fatal(err)
	}
	// initialiaze authService client
	authServiceClient = pb.NewAuthServicesClient(grpcConn)

	// Initialize layers
	userService = services.NewUserServiceImpl(userCollection, authServiceClient, grpcConn, ctx)
	userController = controllers.NewUserControllers(userService)

	// Initialize gin server
	server = gin.Default()
}

// @title Swagger Example API
// @version 1.0
// @description This is a service for user management.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8001
// @BasePath /v1/users
func main() {
	defer mongoClient.Disconnect(ctx)
	defer grpcConn.Close()

	// Define Base path and register routes
	basePath := server.Group("/v1/users")
	userController.RegisterUserRoutes(basePath)

	// The url pointing to API definition
	url := ginSwagger.URL("http://localhost:8001/swagger/doc.json")
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// start server
	log.Fatal(server.Run(":8001"))
}
