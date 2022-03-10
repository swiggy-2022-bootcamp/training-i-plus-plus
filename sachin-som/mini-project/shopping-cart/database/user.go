package database

import "go.mongodb.org/mongo-driver/mongo"

/*
* Function to create collection for user data.
 */
func CreateUserCollection(client *mongo.Client, databaseName string, collectionName string) *mongo.Collection {
	var userCollection *mongo.Collection = client.Database(databaseName).Collection(collectionName)
	return userCollection
}
