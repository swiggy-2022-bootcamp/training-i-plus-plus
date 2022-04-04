package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"cartService/configs"
	"cartService/models"
	"cartService/responses"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	topic         = "PlacedOrder"
	brokerAddress = "localhost:9092"
)

var i=0

var cartCollection *mongo.Collection = configs.GetCollection(configs.DB, "carts")
var validate = validator.New()

func AddItemToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var cart models.Cart
		defer cancel()

		if err := c.BindJSON(&cart); err != nil {
			c.JSON(http.StatusBadRequest, responses.CartResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&cart); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.CartResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		cart.CartId = primitive.NewObjectID()

		result, err := cartCollection.InsertOne(ctx, cart)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.CartResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetItemsfromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var carts []models.Cart
		defer cancel()

		//objId, _ := primitive.ObjectIDFromHex(userId)

		results, err := cartCollection.Find(ctx, bson.M{"userid": userId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleCart models.Cart
			fmt.Println(results)
			if err = results.Decode(&singleCart); err != nil {
				c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			carts = append(carts, singleCart)
		}

		c.JSON(http.StatusOK, responses.CartResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": carts}})
	}
}

func EditItemFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cartId := c.Param("cartId")
		var cart models.Cart
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(cartId)

		//validate the request body
		if err := c.BindJSON(&cart); err != nil {
			c.JSON(http.StatusBadRequest, responses.CartResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&cart); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.CartResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"quantity": cart.Quantity}
		result, err := cartCollection.UpdateOne(ctx, bson.M{"cartid": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated cart details
		var updatedCart models.Cart
		if result.MatchedCount == 1 {
			err := cartCollection.FindOne(ctx, bson.M{"cartid": objId}).Decode(&updatedCart)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.CartResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedCart}})
	}
}

func DeleteItemFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cartId := c.Param("cartId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(cartId)

		result, err := cartCollection.DeleteOne(ctx, bson.M{"cartid": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.CartResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Cart with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.CartResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Cart successfully deleted!"}},
		)
	}
}

func produce(message []byte) {
	ctx := context.Background()
	l := log.New(os.Stdout, "kafka writer: ", 0)
	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger: l,
		RequiredAcks: 1,
	})

	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(strconv.Itoa(i)),
		Value: message,
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
	i++
	
}


func PlaceOrderFromCart() gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var carts []models.Cart
		defer cancel()

		results, err := cartCollection.Find(ctx, bson.M{"userid": userId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleCart models.Cart
			if err = results.Decode(&singleCart); err != nil {
				c.JSON(http.StatusInternalServerError, responses.CartResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}
			str, _ :=json.Marshal(singleCart)
			fmt.Println(singleCart)
			go produce(str)
			carts = append(carts, singleCart)
		}

		c.JSON(http.StatusOK, responses.CartResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": carts}})
	}
}