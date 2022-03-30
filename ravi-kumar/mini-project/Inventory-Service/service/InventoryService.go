package service

import (
	"Inventory-Service/config"
	errors "Inventory-Service/errors"
	kafka "Inventory-Service/kafka"
	mockdata "Inventory-Service/model"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
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

func CreateProduct(body *io.ReadCloser) (result *mongo.InsertOneResult) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	var newProduct mockdata.Product
	json.NewDecoder(*body).Decode(&newProduct)

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	result, _ = productCollection.InsertOne(ctx, newProduct)
	return
}

func GetCatalog() (allProducts []mockdata.Product) {
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

func GetProductById(productId string) (productRetrieved *mockdata.Product, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(productId)

	if err != nil {
		return nil, errors.MalformedIdError()
	}

	result := productCollection.FindOne(ctx, bson.M{"_id": objectId})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		return nil, errors.IdNotFoundError()
	}

	result.Decode(&productRetrieved)
	return
}

func UpdateProductById(productId string, body *io.ReadCloser) (productRetrieved *mockdata.Product, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	var updatedProduct mockdata.Product
	unmarshalErr := json.NewDecoder(*body).Decode(&updatedProduct)
	if unmarshalErr != nil {
		return nil, errors.UnmarshallError()
	}

	return UpdateProductByIdWorker(productId, updatedProduct)
}

func UpdateProductByIdWorker(productId string, updatedProduct mockdata.Product) (productRetrieved *mockdata.Product, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return nil, errors.MalformedIdError()
	}

	result, error := productCollection.UpdateByID(ctx, objectId, bson.M{"$set": updatedProduct})
	if error != nil {
		return nil, errors.InternalServerError()
	}

	if result.MatchedCount == 0 {
		return nil, errors.IdNotFoundError()
	}
	return GetProductById(productId)
}

func DeleteProductbyId(productId string) (successMessage *string, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return nil, errors.MalformedIdError()
	}

	result, error := productCollection.DeleteOne(ctx, bson.M{"_id": objectId})

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

func UpdateProductQuantity(productId string, updateCount int) (quantityAfterUpdation *int, err error) {
	productRetrieved, error := GetProductById(productId)

	if error != nil {
		productError, ok := error.(*errors.ProductError)
		if ok {
			return nil, productError
		} else {
			fmt.Println("productError casting error in UpdateProductQuantity")
			return
		}
	}

	productRetrieved.QuantityLeft += updateCount
	if productRetrieved.QuantityLeft < 0 {
		ctx, _ := context.WithTimeout(context.Background(), time.Minute*10)
		kafka.Produce(ctx, nil, []byte("productId: "+productId+" --- status: out of stock (critical)"))

		return nil, errors.OutOfStockError()
	}

	//if quantity below threshold, notify monitoring service
	if productRetrieved.QuantityLeft < 20 {
		ctx, _ := context.WithTimeout(context.Background(), time.Minute*10)
		kafka.Produce(ctx, nil, []byte("productId: "+productId+" --- status: quantity below threshold ("+strconv.Itoa(productRetrieved.QuantityLeft)+")"))
	}

	_, err = UpdateProductByIdWorker(productId, *productRetrieved)
	if err != nil {
		return nil, err
	}

	return &productRetrieved.QuantityLeft, nil
}
