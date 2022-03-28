package service

import (
	"Inventory-Service/config"
	errors "Inventory-Service/errors"
	mockdata "Inventory-Service/model"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
		return nil, &errors.ProductError{Status: http.StatusBadRequest, ErrorMessage: "Malformed prodcut id"}
	}

	result := productCollection.FindOne(ctx, bson.M{"_id": objectId})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		return nil, &errors.ProductError{Status: http.StatusNotFound, ErrorMessage: "product with given id not found"}
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
		return nil, &errors.ProductError{Status: http.StatusBadRequest, ErrorMessage: "Couldn't unmarshall user body in request"}
	}

	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return nil, &errors.ProductError{Status: http.StatusBadRequest, ErrorMessage: "Malformed product id"}
	}

	result, error := productCollection.UpdateByID(ctx, objectId, bson.M{"$set": updatedProduct})
	if error != nil {
		return nil, &errors.ProductError{Status: http.StatusInternalServerError, ErrorMessage: "Internal server error"}
	}

	if result.MatchedCount == 0 {
		return nil, &errors.ProductError{Status: http.StatusNotFound, ErrorMessage: "product with given id not found"}
	}
	return GetProductById(productId)
}

func DeleteProductbyId(productId string) (successMessage *string, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return nil, &errors.ProductError{Status: http.StatusBadRequest, ErrorMessage: "Malformed product id"}
	}

	result, error := productCollection.DeleteOne(ctx, bson.M{"_id": objectId})

	if error != nil {
		return nil, &errors.ProductError{Status: http.StatusInternalServerError, ErrorMessage: "Internal server error"}
	}

	if result.DeletedCount == 0 {
		return nil, &errors.ProductError{Status: http.StatusNotFound, ErrorMessage: "product with given id not found"}
	}
	msg := "product deleted"
	successMessage = &msg
	return
}
