package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type trainsWithSchedule struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Stations []Station          `json:"stations" bson:"stations"`
	Seats    [][]bool           `json:"seats" bson:"seats"`
}

type Schedule struct {
	Date   primitive.DateTime   `json:"date" bson:"date"`
	Trains []trainsWithSchedule `json:"trains" bson:"trains"`
}
