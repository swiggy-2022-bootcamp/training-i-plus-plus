package main

import (
	"context"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/ListingService/controllers"
	goKafka "github.com/go-kafka-microservice/ListingService/goKafka/consumer"
	"github.com/go-kafka-microservice/ListingService/routes"
	"github.com/go-kafka-microservice/ListingService/services"
)

var (
	err               error
	server            *gin.Engine
	ctx               context.Context
	kafkaConsumer     *kafka.Consumer
	kafkaService      goKafka.GoKafkaServices
	listingService    services.ListingService
	listingController *controllers.ListingController
	listingRouter     *routes.ListingRoutes
)

func init() {
	// Create KafkaConsumer
	kafkaConsumer, err = goKafka.CreateKafkaConsumer(goKafka.ConsumerConfig())
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Kafka Service
	kafkaService = goKafka.NewGokafkaServiceImpl(kafkaConsumer)

	// Initialize Listing Service
	listingService = services.NewListingServiceImpl(kafkaService, ctx)

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

	// Start Server
	log.Fatal(server.Run(":8003"))
}
