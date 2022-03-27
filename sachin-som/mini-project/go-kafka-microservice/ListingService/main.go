package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/ListingService/controllers"
	"github.com/go-kafka-microservice/ListingService/database"
	goKafka "github.com/go-kafka-microservice/ListingService/goKafka/consumer"
	"github.com/go-kafka-microservice/ListingService/routes"
	"github.com/go-kafka-microservice/ListingService/services"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	err               error
	server            *gin.Engine
	kafkaConsumer     *kafka.Reader
	ctx               context.Context
	mongoClient       *mongo.Client
	listingRouter     *routes.ListingRoutes
	productCollection *mongo.Collection
	kafkaService      goKafka.GoKafkaServices
	listingService    services.ListingService
	listingController *controllers.ListingController
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
	kafkaConsumer = goKafka.CreateKafkaConsumer(goKafka.ConsumerConfig())

	// Initialize Kafka Service
	kafkaService = goKafka.NewGokafkaServiceImpl(kafkaConsumer, productCollection, ctx)

	// Initialize Listing Service
	listingService = services.NewListingServiceImpl(kafkaService, productCollection, ctx)

	// Initialize listing Controller
	listingController = controllers.NewListingController(listingService)

	// Initialize listing Router
	listingRouter = routes.NewListingRoutes(listingController)

	// Initialize Server
	server = gin.Default()
}

func main() {

	// Register All routes
	basePath := server.Group("/v1/listing")
	listingRouter.RegisterListingRoutes(basePath)

	// Consume Products from Inventory Service and
	// Store them to ProductDB's product collection
	go kafkaService.StoreProducts("products")

	// Start Server
	log.Fatal(server.Run(":8003"))
}
