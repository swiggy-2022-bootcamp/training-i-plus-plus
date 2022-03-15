package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseConn() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!!!")
	return client
}

var DB *mongo.Client = DatabaseConn()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("maindb").Collection(collectionName)
	return collection
}
