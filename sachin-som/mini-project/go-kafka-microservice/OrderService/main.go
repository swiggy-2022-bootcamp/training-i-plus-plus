package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/OrderService/controllers"
	"github.com/go-kafka-microservice/OrderService/database"
	gokafkaConsumer "github.com/go-kafka-microservice/OrderService/goKafka/consumer"
	"github.com/go-kafka-microservice/OrderService/routes"
	"github.com/go-kafka-microservice/OrderService/services"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	err                  error
	server               *gin.Engine
	ctx                  context.Context
	mongoClient          *mongo.Client
	kafkaConsumer        *kafka.Reader
	orderCollection      *mongo.Collection
	orderRoutes          *routes.OrderRoutes
	orderService         services.OrderServices
	orderControllers     *controllers.OrderControllers
	kafkaConsumerService gokafkaConsumer.GoKafkaServices
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

	// Initialize kafkaConsumerServices
	kafkaConsumer = gokafkaConsumer.CreateKafkaConsumer(gokafkaConsumer.ConsumerConfig())
	kafkaConsumerService = gokafkaConsumer.NewGokafkaServiceImpl(kafkaConsumer, orderCollection, ctx)

	// Initialize Order Service
	orderService = services.NewOrderServiceImpl(orderCollection, ctx)

	// Initialize Order Controllers
	orderControllers = controllers.NewOrderCollection(orderService)

	// Initialize Order Routes
	orderRoutes = routes.NewListingRoutes(orderControllers)

	// Initialize gin server
	server = gin.Default()
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
