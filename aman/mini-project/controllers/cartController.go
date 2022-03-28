package controllers

import (
	"aman-swiggy-mini-project/database"
	"aman-swiggy-mini-project/models"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "cart")

func GetCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		// orderId := c.Param("cart_id")
		// var order models.Cart

		////////
		findOptions := options.Find()
		var results []models.User
		cur, err := userCollection.Find(ctx, bson.M{"user_type": "Buyer"}, findOptions)
		if err != nil {
			fmt.Println(err)
		}
		for cur.Next(ctx) {
			var elem models.User
			err := cur.Decode(&elem)
			if err != nil {
				fmt.Println(err)
			}

			results = append(results, elem)

		}

		if err := cur.Err(); err != nil {
			fmt.Println(err)
		}

		cur.Close(ctx)

		////////

		// err := orderCollection.FindOne(ctx, bson.M{"cart_id": orderId}).Decode(&order)
		defer cancel()
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the orders"})
		// }
		c.JSON(http.StatusOK, results)
	}
}

func CreateCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var order models.Cart

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(order)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		order.ID = primitive.NewObjectID()
		order.Cart_id = order.ID.Hex()

		result, insertErr := orderCollection.InsertOne(ctx, order)

		if insertErr != nil {
			msg := fmt.Sprintf("order item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}
