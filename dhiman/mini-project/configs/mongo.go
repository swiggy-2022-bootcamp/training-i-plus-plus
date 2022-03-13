
package configs

import (
    "context"
    "fmt"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func connectDB() *mongo.Client  {
    client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
    if err != nil {
        log.Fatal(err)
    }
  
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    //ping the database
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB")
    return client
}

//Client instance
var db *mongo.Client = connectDB()

//getting database collections
func getCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    collection := client.Database("golangAPI").Collection(collectionName)
    return collection
}

var UsersCollection *mongo.Collection = getCollection(db, UsersCollectionName())
var MedicinesCollection *mongo.Collection = getCollection(db, MedicinesCollectionName())
