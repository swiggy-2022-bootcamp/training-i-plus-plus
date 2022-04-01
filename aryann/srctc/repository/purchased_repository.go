package repository

import (
	"context"
	"fmt"
	"srctc/database"
	"srctc/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collectionPurchasedName = "purchased"
	collectionPurchased     = new(mongo.Collection)
)

func init() {
	var DB *mongo.Client = database.ConnectDB()
	collectionPurchased = DB.Database("golangAPI").Collection(collectionPurchasedName)
}

type PurchasedRepository struct{}

func (btk PurchasedRepository) Create(newPurchased models.Purchased) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionPurchased.InsertOne(ctx, &newPurchased)
	if err == nil {
		fmt.Println("Inserted a single document: ", result.InsertedID)
	}
	return result.InsertedID, err
}

func (btk PurchasedRepository) Read(objId primitive.ObjectID) (models.Purchased, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var user models.Purchased
	err := collectionPurchased.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	return user, err
}

func (btk PurchasedRepository) Delete(objId primitive.ObjectID) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collectionPurchased.DeleteOne(ctx, bson.M{"_id": objId})
	if err == nil {
		fmt.Println("Updated a single document: ", result.DeletedCount)
	}
	return result.DeletedCount, err
}

func (btk PurchasedRepository) ReadAll() ([]models.Purchased, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	var users []models.Purchased
	defer cancel()

	results, err := collectionPurchased.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer results.Close(ctx)

	for results.Next(ctx) {
		var singlePurchased models.Purchased
		if err = results.Decode(&singlePurchased); err != nil {
			return users, err
		}

		users = append(users, singlePurchased)
	}
	return users, nil
}
