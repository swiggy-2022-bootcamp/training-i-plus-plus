
package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	//"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const (
	// MongoDBUrl is the default mongodb url that will be used to connect to the
	// database.
	dbName="myFirstDatabase"
	
	MongoDBUrl =""
)

func DBintance() *mongo.Client {
	

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv(MongoDBUrl)))

	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Mongo DB")

	return client
}

var Client *mongo.Client = DBintance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	
	var collection *mongo.Collection = client.Database(dbName).Collection(collectionName)
	return collection
}
