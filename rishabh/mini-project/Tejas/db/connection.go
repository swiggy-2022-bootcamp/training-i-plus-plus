package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getConnection() *mongo.Client {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	mongoDBuri := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(mongoDBuri)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MonogoDB Connected")

	return client
}

var MongoClient *mongo.Client = getConnection()

func GetCollection(collectionName string) *mongo.Collection {
	return MongoClient.Database("tejas").Collection(collectionName)
}
