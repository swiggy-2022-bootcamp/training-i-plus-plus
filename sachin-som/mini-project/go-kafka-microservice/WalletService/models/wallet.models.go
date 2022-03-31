package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Wallet struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id" validate:"required"`
	Amount    int                `json:"amount" bson:"amount"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
