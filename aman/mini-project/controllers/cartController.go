package controllers

import (
	"aman-swiggy-mini-project/models"
	"context"
	"fmt"
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

		if err := cur.Err(); err != nil {
			fmt.Println(err)
		}

		var productItem models.Product
		var finalResults []models.Product
		for _, elem := range results {
			productId := elem.Product_id
			err := productCollection.FindOne(ctx, bson.M{"product_id": productId}).Decode(&productItem)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured finding product items"})
			}
			finalResults = append(finalResults, productItem)
		}

		cur.Close(ctx)

		defer cancel()
		c.JSON(http.StatusOK, finalResults)
	}
}
