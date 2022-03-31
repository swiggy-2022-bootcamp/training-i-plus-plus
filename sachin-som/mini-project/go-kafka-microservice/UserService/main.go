package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	pb "github.com/go-kafka-microservice/AuthService/controllers"
	"github.com/go-kafka-microservice/UserService/controllers"
	"github.com/go-kafka-microservice/UserService/database"
	"github.com/go-kafka-microservice/UserService/services"
	"github.com/joho/godotenv"
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

func main() {
	defer mongoClient.Disconnect(ctx)
	defer grpcConn.Close()

	// Define Base path and register routes
	basePath := server.Group("/v1/users")
	userController.RegisterUserRoutes(basePath)

	// start server
	log.Fatal(server.Run(":8001"))
}
