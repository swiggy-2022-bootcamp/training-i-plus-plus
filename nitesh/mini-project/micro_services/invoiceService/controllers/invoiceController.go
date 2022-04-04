package controllers

import (
	"context"
	"invoiceService/database"
	"invoiceService/logger"
	"invoiceService/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var invoiceCollection *mongo.Collection = database.OpenCollection(database.MongoClient, "invoiceDetails")
var log logrus.Logger = *logger.GetLogger()

var request struct {
	Id string
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.WithFields(logrus.Fields{"error": err.Error()}).Error("not able to bind with json")
			return
		}

		var invoice models.Invoice
		err := invoiceCollection.FindOne(ctx, bson.M{"id": request.Id}).Decode(&invoice)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.WithFields(logrus.Fields{"error": err.Error()}).Error("invoice details not found")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"invoiceDetails": invoice,
		})
	}
}
