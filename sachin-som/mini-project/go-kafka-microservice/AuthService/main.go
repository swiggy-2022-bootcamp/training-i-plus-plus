package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/go-kafka-microservice/AuthProto"
	"github.com/go-kafka-microservice/AuthService/controllers"
	"github.com/go-kafka-microservice/AuthService/database"
	"github.com/go-kafka-microservice/AuthService/services"
	"github.com/go-kafka-microservice/AuthService/utils"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	ctx             context.Context
	err             error
	httpListener    net.Listener
	grpcServer      *grpc.Server
	jwtUtils        utils.JWTUtils
	mongoClient     *mongo.Client
	userCollection  *mongo.Collection
	authServices    services.AuthServices
	authControllers *controllers.AuthControllers
)

func init() {
	// Parse .dot file env variables
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error Loading in .env file: ", err.Error())
	}

	// Context Creation
	ctx = context.TODO()
	// Initialize jwt utils service
	jwtUtils = utils.NewJWTUtils()

	// MongoClient Connection
	mongoClient = database.SetUpDatabase(ctx)

	// Create User Collection in MongoDB
	userCollection = mongoClient.Database("UserDB").Collection("users")

	// Initialize authService
	authServices = services.NewAuthServiceImpl(jwtUtils, userCollection, ctx)

	// Initialize (gRPC Service Implementation Server) authControllers
	authControllers = controllers.NewAuthControllers(authServices)

	// Create TCP HTTP connection
	httpListener, err = net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("Error in httpListening (8000): ", err.Error())
	}
}
func main() {
	defer mongoClient.Disconnect(ctx)

	// Create gRPC Server
	grpcServer = grpc.NewServer()

	// Register gRPC Services
	pb.RegisterAuthServicesServer(grpcServer, *authControllers)

	// Start the Server
	if err = grpcServer.Serve(httpListener); err != nil {
		log.Fatal(err)
	}
	fmt.Println("(AuthService): Started Server on Port 8000.")
}
