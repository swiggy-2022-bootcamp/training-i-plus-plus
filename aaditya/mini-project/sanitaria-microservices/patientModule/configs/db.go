package configs

import(
	"log"
	"fmt"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	//initialize client
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMonogoURI()))
	if err != nil {
		log.Fatal(err)
	}

	//Connect to database
	ctx,_ := context.WithTimeout(context.Background(),time.Second * 10)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx,nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

//create instance of client to be used in the application
//Singleton design pattern.
var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("patientDB").Collection(collectionName)
	return collection
}