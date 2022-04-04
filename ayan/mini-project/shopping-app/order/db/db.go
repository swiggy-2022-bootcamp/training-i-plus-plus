package db

import (
	"context"
	"time"

	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/ayan/mini-project/shopping-app/order/utils/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDbClient() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb+srv://ayan-dutta-ipp:ayan-dutta-ipp@cluster0.bh90q.mongodb.net/orderDB?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Fatal("Error while connecting to DB : " + err.Error())
		panic(err)
	} else {
		logger.Info("Database Connected Successfully......")

	}
	return client
}

func Collection(client *mongo.Client, collectionName string) *mongo.Collection {

	return client.Database("orderDB").Collection(collectionName)
}
