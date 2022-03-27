package main

import (
	"context"
	"log"

	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/ListingService/controllers"
	"github.com/go-kafka-microservice/ListingService/database"
	gokafkaConsumer "github.com/go-kafka-microservice/ListingService/goKafka/consumer"
	gokafkaProducer "github.com/go-kafka-microservice/ListingService/goKafka/producer"
	"github.com/go-kafka-microservice/ListingService/routes"
	"github.com/go-kafka-microservice/ListingService/services"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	err                  error
	server               *gin.Engine
	kafkaConsumer        *kafka.Reader
	kafkaProducer        *confluentKafka.Producer
	ctx                  context.Context
	mongoClient          *mongo.Client
	listingRouter        *routes.ListingRoutes
	productCollection    *mongo.Collection
	kafkaConsumerService gokafkaConsumer.GoKafkaServices
	kafkaProducerService gokafkaProducer.GoKafkaServices
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

	// Initialize Listing Service
	listingService = services.NewListingServiceImpl(kafkaConsumerService, kafkaProducerService, productCollection, ctx)

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
	go kafkaConsumerService.StoreProducts("products")

	// Start Server
	log.Fatal(server.Run(":8003"))
}
