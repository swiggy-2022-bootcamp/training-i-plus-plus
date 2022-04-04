package configs

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func ConnectDB() *mongo.Client  {
    client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
    if err != nil {
        zap.L().Fatal(err.Error())
    }
  
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        zap.L().Fatal(err.Error())
    }

    //ping the database
    err = client.Ping(ctx, nil)
    if err != nil {
        zap.L().Fatal(err.Error())
    }
    zap.L().Info("Connected to MongoDB")
    fmt.Println("Connected to MongoDB")
    return client
}

//Client instance
var DB *mongo.Client = ConnectDB()

//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    collection := client.Database("TrainBooking").Collection(collectionName)
    return collection
}