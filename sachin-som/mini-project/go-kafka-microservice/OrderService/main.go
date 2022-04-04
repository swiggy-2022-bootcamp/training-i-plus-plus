package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/InventoryService/middleware"
	"github.com/go-kafka-microservice/OrderService/controllers"
	"github.com/go-kafka-microservice/OrderService/database"
	gokafkaConsumer "github.com/go-kafka-microservice/OrderService/goKafka/consumer"
	"github.com/go-kafka-microservice/OrderService/routes"
	"github.com/go-kafka-microservice/OrderService/services"
	pb "github.com/go-kafka-microservice/WalletProto"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	err                  error
	server               *gin.Engine
	ctx                  context.Context
	mongoClient          *mongo.Client
	kafkaConsumer        *kafka.Reader
	orderCollection      *mongo.Collection
	clientConnInt        *grpc.ClientConn
	orderRoutes          *routes.OrderRoutes
	orderService         services.OrderServices
	orderControllers     *controllers.OrderControllers
	kafkaConsumerService gokafkaConsumer.GoKafkaServices
	walletProtoClient    pb.WalletServiceClient
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
	orderCollection = mongoClient.Database("OrderDB").Collection("orders")

	// gRPC client
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	clientConnInt, err = grpc.Dial("localhost:8005", opts...)
	if err != nil {
		log.Fatal(err)
	}
	walletProtoClient = pb.NewWalletServiceClient(clientConnInt)

	// Initialize kafkaConsumerServices
	kafkaConsumer = gokafkaConsumer.CreateKafkaConsumer(gokafkaConsumer.ConsumerConfig())
	kafkaConsumerService = gokafkaConsumer.NewGokafkaServiceImpl(kafkaConsumer, orderCollection, walletProtoClient, ctx)

	// Initialize Order Service
	orderService = services.NewOrderServiceImpl(orderCollection, ctx)

	// Initialize Order Controllers
	orderControllers = controllers.NewOrderCollection(orderService)

	// Initialize Order Routes
	orderRoutes = routes.NewListingRoutes(orderControllers)

	// Initialize gin server
	server = gin.Default()

	// Don't need color console for logging to a file
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("logger.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// Initialize server
	server = gin.Default()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// gin.DefaultWriter = file writer and os.Stdout
	server.Use(gin.LoggerWithFormatter(middleware.FormatLogger))
}

func main() {
	defer mongoClient.Disconnect(ctx)

	go kafkaConsumerService.StoreOrders("ordered_products")

	// Register Order Routes
	basePath := server.Group("/v1/orders")
	orderRoutes.RegisterOrderRoutes(basePath)

	// start server
	log.Fatal(server.Run(":8004"))
}
