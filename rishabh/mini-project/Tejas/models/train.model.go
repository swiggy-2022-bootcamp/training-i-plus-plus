package models

import (
	"tejas/configs"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Station struct {
	Code          string    `json:"code" bson:"code"`
	ArrivalTime   time.Time `json:"arrival_time" bson:"arrival_time"`
	DepartureTime time.Time `json:"departure_time" bson:"departure_time"`
}

type Train struct {
	Id               int       `json:"_id" bson:"_id"`
	Name             string    `json:"name" bson:"name"`
	Stations         []Station `json:"stations" bson:"stations"`
	PerStationCharge int       `json:"per_station_charge" bson:"per_station_charge"`
}

var TrainCollection *mongo.Collection = configs.GetCollection(configs.DB, "trains")
