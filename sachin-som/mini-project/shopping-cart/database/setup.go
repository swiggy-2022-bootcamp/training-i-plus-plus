package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setUpDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		log.Println("Error: Failed to connect with mongoDB server.")
		return nil
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Error: Failed to connect with mongoDB server.")
		return nil
	}
	fmt.Println("Sucess: Connected with mongoDB server.")
	return client
}
