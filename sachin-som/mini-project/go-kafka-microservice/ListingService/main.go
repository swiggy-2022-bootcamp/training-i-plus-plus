package main

import (
	"context"
	"io"
	"log"
	"os"

	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	pb "github.com/go-kafka-microservice/AuthProto"
	"github.com/go-kafka-microservice/ListingService/controllers"
	"github.com/go-kafka-microservice/ListingService/database"
	gokafkaConsumer "github.com/go-kafka-microservice/ListingService/goKafka/consumer"
	gokafkaProducer "github.com/go-kafka-microservice/ListingService/goKafka/producer"
	"github.com/go-kafka-microservice/ListingService/middleware"
	"github.com/go-kafka-microservice/ListingService/routes"
	"github.com/go-kafka-microservice/ListingService/services"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	err                  error
	server               *gin.Engine
	kafkaConsumer        *kafka.Reader
	mongoClient          *mongo.Client
	grpcConn             *grpc.ClientConn
	ctx                  context.Context
	kafkaProducer        *confluentKafka.Producer
	listingRouter        *routes.ListingRoutes
	productCollection    *mongo.Collection
	authProtoClient      pb.AuthServicesClient
	kafkaConsumerService gokafkaConsumer.GoKafkaServices
	kafkaProducerService gokafkaProducer.GoKafkaServices
	listingMiddleware    *middleware.ListingMiddleware
	listingService       services.ListingService
	listingController    *controllers.ListingController
)

func init() {
	// Load Dotenv file
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error Loading in .env file: ", err.Error())
	}

	ctx = context.TODO()

	// initialize mongo Client
	mongoClient = database.SetUpDatabase(ctx)

	// Create Product collection
	productCollection = mongoClient.Database("ProductDB").Collection("products")

	// Create KafkaConsumer
	kafkaConsumer = gokafkaConsumer.CreateKafkaConsumer(gokafkaConsumer.ConsumerConfig())
	kafkaProducer, _ = gokafkaProducer.CreateProducer(gokafkaProducer.Cfg())

	// Initialize Kafka Service
	kafkaConsumerService = gokafkaConsumer.NewGokafkaServiceImpl(kafkaConsumer, productCollection, ctx)
	kafkaProducerService = gokafkaProducer.NewKafkaProducer(kafkaProducer)

	// Create gRPC Client
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	grpcConn, err = grpc.Dial("localhost:8000", opts...)
	if err != nil {
		log.Fatal(err)
	}

	// initialiaze authService client
	authProtoClient = pb.NewAuthServicesClient(grpcConn)

	// Initialize Listing Service
	listingService = services.NewListingServiceImpl(kafkaConsumerService, kafkaProducerService, productCollection, authProtoClient, ctx)

	// Initialize Listing Middleware
	listingMiddleware = middleware.NewListingMiddleware(listingService)

	// Initialize listing Controller
	listingController = controllers.NewListingController(listingService)

	// Initialize listing Router
	listingRouter = routes.NewListingRoutes(listingController)

	// Initialize Server
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

	// Register Middlewares
	server.Use(listingMiddleware.AuthorizeUser())

	// Register All routes
	basePath := server.Group("/v1/listing")
	listingRouter.RegisterListingRoutes(basePath)

	// Consume Products from Inventory Service and
	// Store them to ProductDB's product collection
	go kafkaConsumerService.StoreProducts("products")

	// Start Server
	log.Fatal(server.Run(":8003"))
}
