package models

import (
	"time"
)

type TrainsWithSchedule struct {
	Id               int       `json:"_id" bson:"_id"`
	Stations         []Station `json:"stations" bson:"stations"`
	Seats            [][]bool  `json:"seats" bson:"seats"`
	PerStationCharge int       `json:"per_station_charge" bson:"per_station_charge"`
}

type Schedule struct {
	Date   time.Time            `json:"date" bson:"date"`
	Trains []TrainsWithSchedule `json:"trains" bson:"trains"`
}
