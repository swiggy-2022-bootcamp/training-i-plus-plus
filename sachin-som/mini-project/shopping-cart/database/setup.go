package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	err         error
	mongoClient *mongo.Client
)

func SetUpDB(ctx context.Context) *mongo.Client {
	mongoCredentials := options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		Username:      os.Getenv("MONGO_USERNAME"),
		Password:      os.Getenv("MONGO_PASSWORD"),
	}
	connString := os.Getenv("MONGO_URI")
	mongoConn := options.Client().ApplyURI(connString).SetAuth(mongoCredentials)
	mongoClient, err = mongo.Connect(ctx, mongoConn)
	if err != nil {
		log.Fatal(err)
	}
	// Check MongoConnection by pinging
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo Connection has beed established.")
	return mongoClient
}
