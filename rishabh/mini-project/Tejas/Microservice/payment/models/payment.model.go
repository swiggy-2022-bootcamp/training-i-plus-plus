package models

import (
	"paymentService/configs"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Payment struct {
	TransactionId string             `json:"transaction_id" bson:"transaction_id"`
	UserId        primitive.ObjectID `json:"user_id" bson:"user_id"`
	Amount        int                `json:"amount" bson:"amount"`
	Status        string             `json:"status" bson:"status"`
}

var PaymentCollection *mongo.Collection = configs.GetCollection(configs.DB, "payments")
