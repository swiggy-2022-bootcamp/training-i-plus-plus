package database

import (
	"context"
	"fmt"
	"medo-healthcare-app/cmd/model"
	"medo-healthcare-app/pkg/env"
	"medo-healthcare-app/pkg/err"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString = env.GoDotEnvVariable("MONGOURL")
var dbName = env.GoDotEnvVariable("DBNAME")
var collectionName = env.GoDotEnvVariable("COLLECTIONNAME")

var collection *mongo.Collection

//Establishing connection with the database
func init() {
	mongoClient, cancel := mongo.Connect(context.TODO(), (options.Client().ApplyURI(connectionString)))
	err.CheckNilErr(cancel)
	fmt.Println("MongoDB - Database Connection Established !")
	collection = mongoClient.Database(dbName).Collection(collectionName)
	fmt.Println("MongoDB - Collection Instance Ready !")
}

//Database Helper Methods

//InsertOne - CREATE
func InsertOne(user model.CoreUserData) {
	insertedData, cancel := collection.InsertOne(context.Background(), user)
	err.CheckNilErr(cancel)
	fmt.Println("Inserted One User with UserID : ", insertedData.InsertedID)
}

//Find - READ
func Find() []primitive.M {
	currentValue, cancel := collection.Find(context.Background(), bson.D{{}})
	err.CheckNilErr(cancel)
	var users []primitive.M
	for currentValue.Next(context.Background()) {
		var user bson.M
		cancel := currentValue.Decode(&user)
		err.CheckNilErr(cancel)
		users = append(users, user)
	}
	defer currentValue.Close(context.Background())
	return users
}

//FindOne - SEARCH
func FindOne(emailAddress string) model.CoreUserData {
	filter := bson.M{"email": emailAddress}
	var user bson.M
	var userStruct model.CoreUserData
	cancel := collection.FindOne(context.Background(), filter).Decode(&user)
	err.CheckNilErr(cancel)
	bsonBytes, _ := bson.Marshal(user)
	bson.Unmarshal(bsonBytes, &userStruct)
	return userStruct
}
