package database

import (
	"context"
	"fmt"
	"medo-healthcare-app/cmd/model"
	"medo-healthcare-app/pkg/env"
	"medo-healthcare-app/pkg/logger"

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
	logger.Error("Error : ", cancel)
	collection = mongoClient.Database(dbName).Collection(collectionName)
	fmt.Println("MongoDB - Collection Instance Ready !")
}

//Database Helper Methods

//InsertOne - CREATE
func InsertOne(user model.CoreUserData) {
	insertedData, cancel := collection.InsertOne(context.Background(), user)
	logger.Error("Error : ", cancel)
	fmt.Println("Inserted One User with UserID : ", insertedData.InsertedID)
}

//Find - READ
func Find() []primitive.M {
	currentValue, cancel := collection.Find(context.Background(), bson.D{{}})
	logger.Error("Error : ", cancel)
	var users []primitive.M
	for currentValue.Next(context.Background()) {
		var user bson.M
		cancel := currentValue.Decode(&user)
		logger.Error("Error : ", cancel)
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
	logger.Error("Error : ", cancel)
	bsonBytes, _ := bson.Marshal(user)
	bson.Unmarshal(bsonBytes, &userStruct)
	return userStruct
}

//UpdateOne - UPDATE
func UpdateOne(username string, valueType string, newValue string) model.CoreUserData {
	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{valueType: newValue}}
	result, cancel := collection.UpdateOne(context.Background(), filter, update)
	logger.Error("Error : ", cancel)
	fmt.Println("Modified Count :", result.ModifiedCount)
	newFilter := bson.M{valueType: newValue}
	var user model.CoreUserData
	cancel1 := collection.FindOne(context.Background(), newFilter).Decode(&user)
	logger.Error("Error : ", cancel1)
	if valueType == "email" {
		if newValue != username {
			return UpdateOne(username, "username", newValue)
		}
	}
	return user
	//return FindOne(username)
}

//DeleteOne - DELETE
func DeleteOne(emailAddress string) primitive.M {
	filter := bson.M{"email": emailAddress}
	var user primitive.M
	cancel := collection.FindOneAndDelete(context.Background(), filter).Decode(&user)
	logger.Error("Error : ", cancel)
	fmt.Println("DELETION SUCCESSFUL")
	return user
}

//DeleteMany - DELETE ALL
func DeleteMany() int64 {
	deletedResult, cancel := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	logger.Error("Error : ", cancel)
	fmt.Println("Users Deleted :", deletedResult.DeletedCount)
	return deletedResult.DeletedCount
}
