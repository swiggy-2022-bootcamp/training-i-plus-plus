package controllers

import (
	"context"
	"net/http"
	"paymentService/dto"
	"paymentService/kafka"
	"paymentService/models"
	"paymentService/services"
	"paymentService/utils"
	"time"

	"github.com/gin-gonic/gin"
)

const requestTimeout = 10 * time.Second

func Payment() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()
		var paymentDTO dto.PaymentDTO
		c.BindJSON(&paymentDTO)

		user, _ := c.MustGet("user_details").(services.SignedDetails)

		payment := models.Payment{
			TransactionId: utils.GenerateRandomHash(10),
			UserId:        user.UserId,
			Amount:        paymentDTO.Amount,
			Status:        "success",
		}
		_, err := models.PaymentCollection.InsertOne(ctx, payment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		go kafka.PaymentDetails(payment)
		c.JSON(http.StatusOK, gin.H{"message": "Payment added successfully", "data": payment})
	}
}
