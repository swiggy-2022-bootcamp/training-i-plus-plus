package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/OrderService/database"
	gokafkaConsumer "github.com/go-kafka-microservice/OrderService/goKafka/consumer"
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
	orderService         services.OrderServices
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
	kafkaConsumerService = gokafkaConsumer.NewGokafkaServiceImpl(kafkaConsumer, orderCollection, ctx)

	// Initialize gin server
	server = gin.Default()
}

func main() {
	defer mongoClient.Disconnect(ctx)

	go kafkaConsumerService.StoreOrders("ordered_products")

	// start server
	log.Fatal(server.Run(":8001"))
}
