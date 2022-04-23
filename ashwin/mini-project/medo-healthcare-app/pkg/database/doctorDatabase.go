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

var doctorDBName = env.GoDotEnvVariable("DOCTORDBNAME")
var doctorCollectionName = env.GoDotEnvVariable("DOCTORCOLLECTIONNAME")

var doctorCollection *mongo.Collection

//Establishing connection with the database
func init() {
	mongoClient, cancel := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	logger.Error("Error : ", cancel)
	doctorCollection = mongoClient.Database(doctorDBName).Collection(doctorCollectionName)
	fmt.Println("MongoDB - Doctor Collection Instance Ready !")
}

//Database Helper Methods

//DocInsertOne - CREATE
func DocInsertOne(user model.CoreUserData) {
	insertedData, cancel := doctorCollection.InsertOne(context.Background(), user)
	logger.Error("Error : ", cancel)
	fmt.Println("Inserted One User with UserID : ", insertedData.InsertedID)
}

//DocFind - READ
func DocFind() []primitive.M {
	currentValue, cancel := doctorCollection.Find(context.Background(), bson.D{{}})
	logger.Error("Error : ", cancel)
	var doctors []primitive.M
	for currentValue.Next(context.Background()) {
		var user bson.M
		cancel := currentValue.Decode(&user)
		logger.Error("Decode Error : ", cancel)
		doctors = append(doctors, user)
	}
	defer currentValue.Close(context.Background())
	return doctors
}

//DocFindOne - SEARCH
func DocFindOne(emailAddress string) model.CoreUserData {
	filter := bson.M{"email": emailAddress}
	var user bson.M
	var userStruct model.CoreUserData
	cancel := doctorCollection.FindOne(context.Background(), filter).Decode(&user)
	logger.Error("Error : ", cancel)
	bsonBytes, _ := bson.Marshal(user)
	bson.Unmarshal(bsonBytes, &userStruct)
	return userStruct
}

//DocUpdateOne - UPDATE
func DocUpdateOne(username string, valueType string, newValue string) model.CoreUserData {
	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{valueType: newValue}}
	result, cancel := doctorCollection.UpdateOne(context.Background(), filter, update)
	logger.Error("Error : ", cancel)
	fmt.Println("Modified Count :", result.ModifiedCount)
	newFilter := bson.M{valueType: newValue}
	var user model.CoreUserData
	cancel1 := doctorCollection.FindOne(context.Background(), newFilter).Decode(&user)
	logger.Error("Error : ", cancel1)
	if valueType == "email" {
		if newValue != username {
			return UpdateOne(username, "username", newValue)
		}
	}
	return user
	//return FindOne(username)
}

//DocDeleteOne - DELETE
func DocDeleteOne(emailAddress string) primitive.M {
	filter := bson.M{"email": emailAddress}
	var user primitive.M
	cancel := doctorCollection.FindOneAndDelete(context.Background(), filter).Decode(&user)
	logger.Error("Error : ", cancel)
	fmt.Println("DELETION SUCCESSFUL")
	return user
}

//DocDeleteMany - DELETE ALL
func DocDeleteMany() int64 {
	deletedResult, cancel := doctorCollection.DeleteMany(context.Background(), bson.D{{}}, nil)
	logger.Error("Error : ", cancel)
	fmt.Println("Users Deleted :", deletedResult.DeletedCount)
	return deletedResult.DeletedCount
}
