package models

import (
	"reservationService/configs"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Reservation struct {
	PNR             primitive.ObjectID `json:"pnr" bson:"pnr"`
	UserId          primitive.ObjectID `json:"user_id" bson:"user_id"`
	TrainId         int                `json:"train_id" bson:"train_id"`
	FromStationCode string             `json:"from_station_code" bson:"from_station_code"`
	ToStationCode   string             `json:"to_station_code" bson:"to_station_code"`
	Date            time.Time          `json:"date" bson:"date"`
	TransactionId   string             `json:"transaction_id" bson:"transaction_id"`
	Status          string             `json:"status" bson:"status"`
	SeatNumber      int                `json:"seat_number" bson:"seat_number"`
}

var ReservationCollection *mongo.Collection = configs.GetCollection(configs.DB, "reservations")
