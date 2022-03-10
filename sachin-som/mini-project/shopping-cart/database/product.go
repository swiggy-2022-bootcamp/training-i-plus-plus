package database

import "go.mongodb.org/mongo-driver/mongo"

/*
* Function to create collection for product data.
 */
func CreateProductCollection(client *mongo.Client, databaseName string, collectionName string) *mongo.Collection {
	var productCollection *mongo.Collection = client.Database(databaseName).Collection(collectionName)
	return productCollection
}
