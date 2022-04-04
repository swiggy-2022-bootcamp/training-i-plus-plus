package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"orderService/configs"
	"orderService/models"
	"orderService/responses"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = configs.GetCollection(configs.DB, "orders")
const (
	brokerAddress = "localhost:9092"
)

func ProcessOrder() {
	
	ctx := context.Background()
	l := log.New(os.Stdout, "kafka reader: ", 0)
	l.Println("Started listening for new orders")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   "PlacedOrder",
		GroupID: "Order-process",
		StartOffset: kafka.FirstOffset,
		Logger: l,
		MaxWait: 60 * time.Second,
	})
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Println("received: ", string(msg.Value))
	
		dbctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			var order models.Order
		
			validationErr := json.Unmarshal((msg.Value), &order)
			if validationErr != nil {
				log.Println("Order has insufficient info ",validationErr)
				panic(err)
			}

			order.OrderId = primitive.NewObjectID()
			order.Status="Order Placed"
			order.Total=order.Price*order.Quantity
			result, err := orderCollection.InsertOne(dbctx, order)
			if err != nil {
				log.Println("Insert into Order DB failed ",validationErr)
				panic(err)
			}
			log.Println("Order proccessed and inserted into DB",result)
	}
}

func UpdateOrderStatus() {
	
	ctx := context.Background()
	l := log.New(os.Stdout, "kafka reader: ", 0)
	l.Println("Started listening for orders status updates")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   "UpdateOrderStatus",
		GroupID: "Order-status",
		StartOffset: kafka.FirstOffset,
		Logger: l,
		MaxWait: 60 * time.Second,
	})
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Println("received: ", string(msg.Value))
	
		var order models.Order
		validationErr := json.Unmarshal((msg.Value), &order)
		if validationErr != nil {
			log.Println("Order has insufficient info ",validationErr)
			panic(err)
		}
		fmt.Println("recahce")
		querydbctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		orderId := order.OrderId
		defer cancel()
				
		result,queryErr := orderCollection.UpdateOne(querydbctx, bson.M{"orderid": orderId},bson.M{"$set":bson.M{"status":order.Status}})
		
		if queryErr != nil {
			log.Println("Order doesnt exist ",queryErr)
			panic(queryErr)
		}

		log.Println("Order status updated",result)
	}
}


func GetAllUserOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var orders []models.Order
		defer cancel()

		//objId, _ := primitive.ObjectIDFromHex(userId)

		results, err := orderCollection.Find(ctx, bson.M{"userid": userId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.OrderResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleOrder models.Order
			if err = results.Decode(&singleOrder); err != nil {
				c.JSON(http.StatusInternalServerError, responses.OrderResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			orders = append(orders, singleOrder)
		}

		c.JSON(http.StatusOK, responses.OrderResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": orders}})
	}
}

func GetAllSellerOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		sellerId := c.Param("sellerId")
		var orders []models.Order
		defer cancel()

		//objId, _ := primitive.ObjectIDFromHex(userId)

		results, err := orderCollection.Find(ctx, bson.M{"sellerid": sellerId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.OrderResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleOrder models.Order
			if err = results.Decode(&singleOrder); err != nil {
				c.JSON(http.StatusInternalServerError, responses.OrderResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			orders = append(orders, singleOrder)
		}

		c.JSON(http.StatusOK, responses.OrderResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": orders}})
	}
}

