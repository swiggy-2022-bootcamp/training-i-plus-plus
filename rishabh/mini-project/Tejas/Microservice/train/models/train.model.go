package models

import (
	"trainService/configs"

	"go.mongodb.org/mongo-driver/mongo"
)

type Station struct {
	Code string `json:"code" bson:"code"`
}

type Train struct {
	Id               int       `json:"_id" bson:"_id"`
	Name             string    `json:"name" bson:"name"`
	Stations         []Station `json:"stations" bson:"stations"`
	PerStationCharge int       `json:"per_station_charge" bson:"per_station_charge"`
	Status           string    `json:"status" bson:"status"`
}

var TrainCollection *mongo.Collection = configs.GetCollection(configs.DB, "trains")
