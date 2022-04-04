package models

import (
	"payment/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Payment struct {
	TransactionId primitive.ObjectID `json:"transaction_id" bson:"transaction_id"`
	UserName      primitive.ObjectID `json:"username" bson:"username"`
	Amount        float64            `json:"amount" bson:"amount"`
	Status        string             `json:"status" bson:"status"`
}

type PaymentDTO struct {
	UserName primitive.ObjectID `json:"username"`
	Amount   float64            `json:"amount"`
}

type BookingPaymentInfo struct {
	PNR    float64 `json:"pnr"`
	Amount float64 `json:"amount"`
}

var PaymentCollection *mongo.Collection = database.GetCollection(database.DB, "payments")
