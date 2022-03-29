package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-kafka-microservice/AuthService/database"
	"github.com/go-kafka-microservice/AuthService/services"
	"github.com/go-kafka-microservice/AuthService/utils"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ctx            context.Context
	err            error
	jwtUtils       utils.JWTUtils
	mongoClient    *mongo.Client
	userCollection *mongo.Collection
	authServices   services.AuthServices
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
}
func main() {
	fmt.Println("Auth Service.")
}
