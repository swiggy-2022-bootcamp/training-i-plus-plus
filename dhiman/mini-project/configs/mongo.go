package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(EnvMongoURI()).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		err = client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		} else {
			client, err = mongo.Connect(ctx, clientOptions)
		}
	}
	if err != nil {
		log.Fatal(err)
	}

	// Try pinging the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

//Client instance
var db *mongo.Client = ConnectDB()

//getting database collections
func getCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("healthDB").Collection(collectionName)
	return collection
}

var UsersCollection *mongo.Collection = getCollection(db, UsersCollectionName())
var MedicinesCollection *mongo.Collection = getCollection(db, MedicinesCollectionName())
