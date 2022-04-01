package service

import (
	"Order-Service/config"
	errors "Order-Service/errors"
	"Order-Service/kafka"
	"Order-Service/middleware"
	mockdata "Order-Service/model"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var mongoURL string = config.MONGO_URL
var orderCollection *mongo.Collection

func init() {
	// Initialize a new mongo client with options
	client, _ = mongo.NewClient(options.Client().ApplyURI(mongoURL))
	orderCollection = client.Database("swiggy_mini").Collection("orders")
}

func PlaceOrder(body *io.ReadCloser) (result *mongo.InsertOneResult, err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	var orderPlaced mockdata.Order
	json.NewDecoder(*body).Decode(&orderPlaced)

	userId := orderPlaced.UserId
	if !IsValidUser(userId) {
		return nil, errors.UserNotFoundError()
	}

	success, errorResponse, errorProductIndex := UpdateProductQuantity(userId, orderPlaced.Items, -1)
	if !success {
		errorMessage := ReadCloserToString(errorResponse.Body) + ". Product Id: " + orderPlaced.Items[*errorProductIndex] + " (Order rolled back)"
		return nil, &errors.OrderError{Status: http.StatusBadRequest, ErrorMessage: errorMessage}
	}

	orderPlaced.OrderDate = time.Now()
	orderPlaced.DeliveryDate = orderPlaced.OrderDate.AddDate(0, 0, 6)
	orderPlaced.Status = "confirmed"

	ctx, _ = context.WithTimeout(context.Background(), time.Second*10)
	result, _ = orderCollection.InsertOne(ctx, orderPlaced)
	orderId := result.InsertedID.(primitive.ObjectID).Hex()

	ctx, _ = context.WithTimeout(context.Background(), time.Minute*10)
	kafka.Produce(ctx, nil, []byte("orderId: "+orderId+" --- status: "+orderPlaced.Status))

	return
}

func GetOrders(userId string) (orders []mockdata.Order, err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	//TODO: check if user exists
	if !IsValidUser(userId) {
		return nil, errors.UserNotFoundError()
	}

	ctx, _ = context.WithTimeout(context.Background(), time.Second*10)
	cursor, _ := orderCollection.Find(ctx, bson.M{"userid": userId})

	for cursor.Next(ctx) {
		var order mockdata.Order
		cursor.Decode(&order)
		orders = append(orders, order)
	}
	return
}

func OrderPayment(orderId string) (successMessage *string, err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return nil, errors.MalformedIdError()
	}

	result := orderCollection.FindOne(ctx, bson.M{"_id": objectId})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		return nil, errors.IdNotFoundError()
	}

	var order mockdata.Order
	result.Decode(&order)

	if order.Status == "payment done" || order.Status == "delivered" {
		return nil, errors.OrderAlreadyPaidForError()
	}

	order.Status = "payment done"

	_, error := orderCollection.UpdateByID(ctx, objectId, bson.M{"$set": order})

	if error != nil {
		return nil, errors.InternalServerError()
	}

	ctx, _ = context.WithTimeout(context.Background(), time.Minute*10)
	kafka.Produce(ctx, nil, []byte("orderId: "+orderId+" --- status: "+order.Status))

	str := "order payment successful"
	successMessage = &str
	return
}

func DeliverOrder(orderId string) (successMessage *string, err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client.Connect(ctx)

	//convert userId string to objectId type
	objectId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return nil, errors.MalformedIdError()
	}

	result := orderCollection.FindOne(ctx, bson.M{"_id": objectId})

	if result.Err() != nil && result.Err() == mongo.ErrNoDocuments {
		return nil, errors.IdNotFoundError()
	}

	var order mockdata.Order
	result.Decode(&order)

	if order.Status == "confirmed" {
		return nil, errors.PaymentIncompleteError()
	}

	if order.Status == "delivered" {
		return nil, errors.OrderAlreadyDeliveredError()
	}

	order.Status = "delivered"

	_, error := orderCollection.UpdateByID(ctx, objectId, bson.M{"$set": order})
	if error != nil {
		return nil, errors.InternalServerError()
	}

	ctx, _ = context.WithTimeout(context.Background(), time.Minute*10)
	kafka.Produce(ctx, nil, []byte("orderId: "+orderId+" --- status: "+order.Status))

	str := "order delivered"
	successMessage = &str
	return
}

func IsValidUser(userId string) bool {
	jwtToken, _ := middleware.GenerateJWT(userId, mockdata.Admin)
	url := "http://localhost:5004/users/"

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + jwtToken

	// Create a new request using http
	req, _ := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}

func UpdateProductQuantity(userId string, productIds []string, quantity int) (success bool, errorResponse *http.Response, errorProductIndex *int) {
	jwtToken, _ := middleware.GenerateJWT(userId, mockdata.Admin)

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + jwtToken

	for index, productId := range productIds {
		url := "http://localhost:5002/catalog/" + productId + "/" + strconv.Itoa(quantity)

		// Create a new request using http
		req, _ := http.NewRequest("POST", url, nil)

		// add authorization header to the req
		req.Header.Add("Authorization", bearer)

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			log.Fatalln(err)
		}

		if resp.StatusCode != http.StatusOK {
			//roll back
			UpdateProductQuantity(userId, productIds[:index], 1)
			return false, resp, &index
		}
	}
	return true, nil, nil
}

func ReadCloserToString(body io.ReadCloser) (message string) {
	json.NewDecoder(body).Decode(&message)
	return
}
