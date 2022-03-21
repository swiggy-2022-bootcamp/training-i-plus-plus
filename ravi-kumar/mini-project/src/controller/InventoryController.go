package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	mockdata "src/model"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	// Initialize a new mongo client with options
	client, _ = mongo.NewClient(options.Client().ApplyURI(mongoURL))
}

func CreateProduct(res http.ResponseWriter, req *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	var newProduct mockdata.Product
	json.NewDecoder(req.Body).Decode(&newProduct)
	collection := client.Database("swiggy_mini").Collection("product")

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, newProduct)

	res.Header().Add("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}

func GetCatalog(res http.ResponseWriter, req *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	var allProducts []mockdata.Product
	collection := client.Database("swiggy_mini").Collection("product")

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	for cursor.Next(ctx) {
		var product mockdata.Product
		cursor.Decode(&product)
		allProducts = append(allProducts, product)
	}

	res.Header().Add("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(allProducts)
}

func GetProductById(res http.ResponseWriter, req *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	//storing route variables for a request
	variables := mux.Vars(req)
	var productId string = variables["productId"]
	var productRetrieved mockdata.Product

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("swiggy_mini").Collection("product")

	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode("Malformed prodcut id")
		return
	}

	result := collection.FindOne(ctx, bson.M{"_id": objectId})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode("product with given id not found")
		return
	}

	result.Decode(&productRetrieved)

	res.Header().Add("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(productRetrieved)
}

func UpdateProductById(res http.ResponseWriter, req *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	//storing route variables for a request
	variables := mux.Vars(req)
	productId := variables["productId"]

	reqBody, _ := ioutil.ReadAll(req.Body)

	updatedProduct := &mockdata.Product{}
	unmarshalErr := json.Unmarshal(reqBody, updatedProduct)
	if unmarshalErr != nil {
		log.Print("Couldn't unmarshall user body in request")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("swiggy_mini").Collection("product")

	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode("Malformed product id")
		return
	}

	result, error := collection.UpdateByID(ctx, objectId, bson.M{"$set": updatedProduct})
	if error != nil {
		fmt.Println(error)
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode("Internal server error")
		return
	}
	if result.MatchedCount == 0 {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode("product with given id not found")
		return
	}

	res.Header().Add("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode("product updated")
}

func DeleteProductbyId(res http.ResponseWriter, req *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = client.Connect(ctx)

	variables := mux.Vars(req)
	productId := variables["productId"]

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("swiggy_mini").Collection("product")

	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode("Malformed product id")
		return
	}

	result, error := collection.DeleteOne(ctx, bson.M{"_id": objectId})

	if error != nil {
		fmt.Println(error)
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode("Internal server error")
		return
	}

	if result.DeletedCount == 0 {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode("product with given id not found")
		return
	}

	res.Header().Add("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode("product deleted")
}

// func GetCatalog(res http.ResponseWriter, req *http.Request) {
// 	res.Header().Add("Content-type", "application/json")
// 	res.WriteHeader(http.StatusOK)
// 	json.NewEncoder(res).Encode(mockdata.GetProductCatalog())
// }

// func GetProductById(res http.ResponseWriter, req *http.Request) {
// 	//storing route variables for a request
// 	variables := mux.Vars(req)
// 	productId := variables["productId"]

// 	for _, productDetail := range mockdata.GetProductCatalog() {
// 		productIdInt, _ := strconv.Atoi(productId)
// 		if productDetail.Id == productIdInt {
// 			res.Header().Add("Content-type", "application/json")
// 			res.WriteHeader(http.StatusOK)
// 			json.NewEncoder(res).Encode(productDetail)
// 			return
// 		}
// 	}

// 	res.Header().Add("Content-type", "application/json")
// 	res.WriteHeader(http.StatusNotFound)
// 	json.NewEncoder(res).Encode("product with given id not found")
// }

// func DeleteProductbyId(res http.ResponseWriter, req *http.Request) {
// 	variables := mux.Vars(req)
// 	productId := variables["productId"]

// 	for index, productDetail := range mockdata.GetProductCatalog() {
// 		productIdInt, _ := strconv.Atoi(productId)
// 		if productDetail.Id == productIdInt {
// 			//to mockdata.Catalog[:index], append everything beyond index (excluding index ofcourse).
// 			//since append() takes "elements" to append as 2nd param, use "..." to lay out elements of slice as independent elements
// 			mockdata.Catalog = append(mockdata.Catalog[:index], mockdata.Catalog[index+1:]...)

// 			res.Header().Add("Content-type", "application/json")
// 			res.WriteHeader(http.StatusOK)
// 			json.NewEncoder(res).Encode("product deleted")
// 			return
// 		}
// 	}

// 	res.Header().Add("Content-type", "application/json")
// 	res.WriteHeader(http.StatusNotFound)
// 	json.NewEncoder(res).Encode("product with given id not found")
// }

// func PrintCatalog() {
// 	catalog := mockdata.GetProductCatalog()
// 	for _, product := range catalog {
// 		PrintProduct(&product)
// 	}
// }

// func PrintProduct(product *mockdata.Product) {
// 	fmt.Println("\nname: ", product.Name,
// 		"\nprice: ", product.Price,
// 		"\ndescription: ", product.Description,
// 		"\nseller: ", product.Seller,
// 		"\nrating: ", product.Rating,
// 		"\nreview: ", strings.Join(product.Review, ", "))
// }
