package configs

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectDb() *mongo.Client {

	clientOptions := options.Client().
		ApplyURI(EnvMongoURI())
		
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	return client
}

var DB *mongo.Client = ConnectDb()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("irctc").Collection(collectionName)
	return collection
}
