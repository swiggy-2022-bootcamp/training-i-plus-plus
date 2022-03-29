package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reservation struct {
	PNR               primitive.ObjectID `json:"pnr" bson:"pnr"`
	User_id           primitive.ObjectID `json:"user_id" bson:"user_id"`
	Train_id          int                `json:"train_id" bson:"train_id"`
	From_Station_code string             `json:"from_station_code" bson:"from_station_code"`
	To_Station_code   string             `json:"to_station_code" bson:"to_station_code"`
	Date              time.Time          `json:"date" bson:"date"`
	Transaction_id    string             `json:"transaction_id" bson:"transaction_id"`
	Status            string             `json:"status" bson:"status"`
}
