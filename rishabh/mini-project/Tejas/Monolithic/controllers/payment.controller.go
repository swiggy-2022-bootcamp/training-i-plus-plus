package controllers

import (
	"context"
	"tejas/models"
	"tejas/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Payment(userId primitive.ObjectID, amount int) (models.Payment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	payment := models.Payment{
		Transaction_id: utils.GenerateRandomHash(10),
		User_id:        userId,
		Amount:         amount,
		Status:         "success",
	}
	_, err := models.PaymentCollection.InsertOne(ctx, payment)
	if err != nil {
		return models.Payment{}, err
	}
	return payment, nil
}
