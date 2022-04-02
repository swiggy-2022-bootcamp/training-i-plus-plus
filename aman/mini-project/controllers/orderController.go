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

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "orders")

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		userID := c.Request.Header.Get("id")
		var order models.Order

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if order.Payment_method == "COD" {
			order.Payment_status = "UNPAID"
		} else if order.Payment_method == "ONLINE" {
			order.Payment_status = "PAID"
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Valid Payment Methods": "ONLINE, COD", "error": "Please choose valid payment method"})
			return
		}

		findOptions := options.Find()
		var results []models.CartItem
		cur, err := cartItemCollection.Find(ctx, bson.M{"user_id": userID}, findOptions)
		if err != nil {
			fmt.Println(err)
		}
		for cur.Next(ctx) {
			var elem models.CartItem
			err := cur.Decode(&elem)
			if err != nil {
				fmt.Println(err)
			}
			results = append(results, elem)
		}
		fmt.Println(results)
		if err := cur.Err(); err != nil {
			fmt.Println(err)
		}
		totalCartValue := float64(0)
		var productItem models.Product
		var finalResults []interface{}
		for _, elem := range results {
			productId := elem.Product_id
			err := productCollection.FindOne(ctx, bson.M{"product_id": productId}).Decode(&productItem)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured finding product items"})
			}
			totalCartValue += float64(*elem.Quantity) * float64(*productItem.Price)
			finalResults = append(finalResults, gin.H{"Name": productItem.Name, "Price": productItem.Price, "Quantity": elem.Quantity})
		}

		order.Items = finalResults
		order.ID = primitive.NewObjectID()
		order.Order_id = order.ID.Hex()
		order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.User_id = userID
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

func CancelOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		cartId := c.Request.Header.Get("id")
		orderId := c.Param("order_id")
		result, insertErr := orderCollection.DeleteOne(ctx, bson.M{"user_id": cartId, "order_id": orderId})

		if insertErr != nil {
			msg := fmt.Sprintf("order was not canceled")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func PaidOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		userType := c.Request.Header.Get("type")
		orderId := c.Param("order_id")

		if userType == "Buyer" {
			c.JSON(http.StatusAccepted, gin.H{"error": "Users not allowed to update status"})
			return
		} else if userType == "Seller" {
			result, insertErr := orderCollection.UpdateOne(ctx, bson.M{"order_id": orderId}, bson.D{{"$set",
				bson.D{
					{"payment_status", "PAID"},
				},
			}})

			if insertErr != nil {
				msg := fmt.Sprintf("order item was not updated")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}

			defer cancel()
			c.JSON(http.StatusOK, result)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unauthorized Request. Please login"})
			return
		}

	}
}
