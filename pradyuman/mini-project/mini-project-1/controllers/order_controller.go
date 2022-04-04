package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mini-project/configs"
	"mini-project/models"
	"mini-project/responses"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = configs.GetCollection(configs.DB, "order")

func ProcessOrder() {
	
	ctx := context.Background()
	l := log.New(os.Stdout, "kafka reader: ", 0)
	l.Println("Started listening for new orders")
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "Order-process",
		StartOffset: kafka.FirstOffset,
		Logger: l,
		MaxWait: 30 * time.Second,
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

			result, err := orderCollection.InsertOne(dbctx, order)
			if err != nil {
				log.Println("Insert into Order DB failed ",validationErr)
				panic(err)
			}
			log.Println("Order proccessed and inserted into DB",result)
	}
}

func GetAllOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var orders []models.Order
		defer cancel()

		//objId, _ := primitive.ObjectIDFromHex(userId)

		results, err := orderCollection.Find(ctx, bson.M{"userid": userId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleOrder models.Order
			fmt.Println(results)
			if err = results.Decode(&singleOrder); err != nil {
				c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			orders = append(orders, singleOrder)
		}

		c.JSON(http.StatusOK, responses.CartResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": orders}})
	}
}