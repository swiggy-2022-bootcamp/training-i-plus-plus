package service

import (
	"context"
	"net/http"
	kafka "payment/kafka"
	model "payment/model"

	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const requestTimeout = 10 * time.Second

func Payment() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()
		var paymentDTO model.PaymentDTO
		c.BindJSON(&paymentDTO)

		payment := model.Payment{
			TransactionId: primitive.NewObjectID(),
			UserName:      paymentDTO.UserName,
			Amount:        paymentDTO.Amount,
			Status:        "success",
		}
		_, err := model.PaymentCollection.InsertOne(ctx, payment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		go kafka.WriteMsgToKafka("payment", payment)
		c.JSON(http.StatusOK, gin.H{"message": "Payment added successfully", "data": payment})
	}
}
