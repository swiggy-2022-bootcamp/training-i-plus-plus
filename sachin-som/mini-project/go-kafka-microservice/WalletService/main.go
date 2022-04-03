package main

import (
	"context"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	pb "github.com/go-kafka-microservice/WalletProto"
	"github.com/go-kafka-microservice/WalletService/controllers"
	"github.com/go-kafka-microservice/WalletService/database"
	"github.com/go-kafka-microservice/WalletService/services"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	server *gin.Engine
	err    error
	ctx    context.Context

	lis              net.Listener
	grpcServer       *grpc.Server
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

	// Create TCP HTTP connection
	lis, err = net.Listen("tcp", ":8005")
	if err != nil {
		log.Fatal("Error in httpListening (8000): ", err.Error())
	}
}

func main() {
	defer mongoClient.Disconnect(ctx)

	base := server.Group("/v1/wallet")
	walletController.RegisterWalletRoutes(base)

	// create grpc server
	grpcServer = grpc.NewServer()

	// Register gRPC Services
	pb.RegisterWalletServiceServer(grpcServer, walletController)

	// Run grpc server on a seprate goRoutine
	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	// Running gin engine on main routine
	log.Fatal(server.Run(":8006"))
}
