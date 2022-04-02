package controllers

import (
	"aman-swiggy-mini-project/database"
	"aman-swiggy-mini-project/models"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cartItemCollection *mongo.Collection = database.OpenCollection(database.Client, "cartItem")

func AddCartItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var item models.CartItem

		if err := c.BindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		item.ID = primitive.NewObjectID()
		result, insertErr := cartItemCollection.InsertOne(ctx, item)

		if insertErr != nil {
			msg := fmt.Sprintf("order item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func RemoveCartItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		cartId := c.Request.Header.Get("id")
		productId := c.Param("product_id")
		result, insertErr := cartItemCollection.DeleteMany(ctx, bson.M{"user_id": cartId, "product_id": productId})

		if insertErr != nil {
			msg := fmt.Sprintf("order item was not deleted")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func UpdateCartItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		cartId := c.Request.Header.Get("id")
		productId := c.Param("product_id")
		var bMap map[string]interface{}

		body := c.Request.Body
		b, err := io.ReadAll(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		json.Unmarshal([]byte(b), &bMap)
		updatedQuantity := bMap["HI"]
		result, insertErr := cartItemCollection.UpdateOne(ctx, bson.M{"user_id": cartId, "product_id": productId}, bson.D{{"$set",
			bson.D{
				{"quantity", updatedQuantity},
			},
		}})

		if insertErr != nil {
			msg := fmt.Sprintf("order item was not updated")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}
