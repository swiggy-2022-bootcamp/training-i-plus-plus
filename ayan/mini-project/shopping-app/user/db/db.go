package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDbClient() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb+srv://ayan-dutta-ipp:ayan-dutta-ipp@cluster0.bh90q.mongodb.net/userDB?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Database Connected Successfully......")

	}
	return client
}

func Collection(client *mongo.Client, collectionName string) *mongo.Collection {

	return client.Database("userDB").Collection(collectionName)
}
