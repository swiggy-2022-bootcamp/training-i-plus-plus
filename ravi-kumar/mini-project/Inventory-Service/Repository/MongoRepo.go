package repository

import (
	"Inventory-Service/config"
	errors "Inventory-Service/errors"
	mockdata "Inventory-Service/model"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var mongoURL string = config.MONGO_URL
var productCollection *mongo.Collection

func init() {
	// Initialize a new mongo client with options
	client, _ = mongo.NewClient(options.Client().ApplyURI(mongoURL))
	productCollection = client.Database("swiggy_mini").Collection("product")
}

type IMongoDAO interface {
	MongoCreateProduct(newProduct mockdata.Product) string
	MongoGetCatalog() (allProducts []mockdata.Product)
	MongoGetProductById(productId primitive.ObjectID) (productRetrieved *mockdata.Product, err error)
	MongoUpdateProductById(productId primitive.ObjectID, updatedProduct mockdata.Product) (productRetrieved *mockdata.Product, err error)
	MongoDeleteProductById(productId primitive.ObjectID) (successMessage *string, err error)
}

type MongoDAO struct {
}

func (dao *MongoDAO) MongoCreateProduct(newProduct mockdata.Product) string {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	result, _ := productCollection.InsertOne(ctx, newProduct)
	return result.InsertedID.(primitive.ObjectID).Hex()
}

func (dao *MongoDAO) MongoGetCatalog() (allProducts []mockdata.Product) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := productCollection.Find(ctx, bson.M{})

	if err != nil {
		fmt.Println(err)
	}

	for cursor.Next(ctx) {
		var product mockdata.Product
		cursor.Decode(&product)
		allProducts = append(allProducts, product)
	}
	return
}

func (dao *MongoDAO) MongoGetProductById(productId primitive.ObjectID) (productRetrieved *mockdata.Product, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	result := productCollection.FindOne(ctx, bson.M{"_id": productId})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		return nil, errors.IdNotFoundError()
	}

	result.Decode(&productRetrieved)
	return
}

func (dao *MongoDAO) MongoUpdateProductById(productId primitive.ObjectID, updatedProduct mockdata.Product) (productRetrieved *mockdata.Product, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	result, error := productCollection.UpdateByID(ctx, productId, bson.M{"$set": updatedProduct})
	if error != nil {
		return nil, errors.InternalServerError()
	}

	if result.MatchedCount == 0 {
		return nil, errors.IdNotFoundError()
	}
	return dao.MongoGetProductById(productId)
}

func (dao *MongoDAO) MongoDeleteProductById(productId primitive.ObjectID) (successMessage *string, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	result, error := productCollection.DeleteOne(ctx, bson.M{"_id": productId})

	if error != nil {
		return nil, errors.InternalServerError()
	}

	if result.DeletedCount == 0 {
		return nil, errors.IdNotFoundError()
	}
	msg := "product deleted"
	successMessage = &msg
	return
}
