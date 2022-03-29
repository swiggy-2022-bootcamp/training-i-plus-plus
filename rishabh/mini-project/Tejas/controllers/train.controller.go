package controllers

import (
	"context"
	"net/http"
	"strconv"
	"tejas/configs"
	"tejas/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var trainCollection *mongo.Collection = configs.GetCollection(configs.DB, "trains")

func AddTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()

		var train models.Train
		c.BindJSON(&train)
		fillTrainDefaults(&train)

		result, err := trainCollection.InsertOne(ctx, train)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"result": result})
	}
}

func RemoveTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()

		var id, err = strconv.Atoi(c.Query("id"))
		result, err := trainCollection.DeleteOne(ctx, bson.M{"_id": id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if result.DeletedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Train with this id doesn't exists "})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": result})
		}
	}
}

func fillTrainDefaults(train *models.Train) {
	if train.PerStationCharge == 0 {
		train.PerStationCharge = 500
	}
}
