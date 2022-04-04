package controllers

import (
	"aman-swiggy-mini-project/logger"
	"aman-swiggy-mini-project/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		cartId := c.Param("cart_id")

		findOptions := options.Find()
		var results []models.CartItem
		cur, err := cartItemCollection.Find(ctx, bson.M{"user_id": cartId}, findOptions)
		if err != nil {
			logger.Log.Println(err)
		}
		for cur.Next(ctx) {
			var elem models.CartItem
			err := cur.Decode(&elem)
			if err != nil {
				logger.Log.Println(err)
			}
			results = append(results, elem)
		}
		if err := cur.Err(); err != nil {
			logger.Log.Println(err)
		}
		totalCartValue := float64(0)
		var productItem models.Product
		var finalResults []interface{}
		for _, elem := range results {
			productId := elem.Product_id
			err := productCollection.FindOne(ctx, bson.M{"product_id": productId}).Decode(&productItem)
			if err != nil {
				logger.Log.Println("Error occured while finding product items")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while finding product items"})
			}
			totalCartValue += float64(*elem.Quantity) * float64(*productItem.Price)
			finalResults = append(finalResults, gin.H{"Name": productItem.Name, "Price": productItem.Price, "Quantity": elem.Quantity})
		}
		cur.Close(ctx)

		defer cancel()
		logger.Log.Println("Cart accessed")
		c.JSON(http.StatusOK, gin.H{"Cart Items": finalResults, "Total Cart Value": totalCartValue})
	}
}
