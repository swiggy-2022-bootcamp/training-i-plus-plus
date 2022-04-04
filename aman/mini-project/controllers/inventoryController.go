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

func GetInventory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		inventoryId := c.Param("inventory_id")

		findOptions := options.Find()
		var results []interface{}
		cur, err := productCollection.Find(ctx, bson.M{"seller_id": inventoryId}, findOptions)
		if err != nil {
			logger.Log.Println(err)
		}
		for cur.Next(ctx) {
			var elem models.Product
			err := cur.Decode(&elem)
			if err != nil {
				logger.Log.Println(err)
			}
			results = append(results, gin.H{"Product ID": elem.Product_id, "Name": elem.Name, "Price": elem.Price, "Units in Stock": elem.Stock_units})
		}
		if err := cur.Err(); err != nil {
			logger.Log.Println(err)
		}

		cur.Close(ctx)
		logger.Log.Println("Inventory accessed")
		defer cancel()
		c.JSON(http.StatusOK, gin.H{"Inventory": results})
	}
}
