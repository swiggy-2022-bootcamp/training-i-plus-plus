package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Train struct {
	ID                 primitive.ObjectID `bson:"_id"`
	TrainNumber        string             `json:"train_number"`
	TrainName          string             `json:"train_name"`
	DepartureStation   string             `json:"departure_station"`
	ArrivalStation     string             `json:"arrival_station"`
	DepartureDate      string             `json:"departure_date"`
	TotalSeatCount     int                `json:"total_seat_count"`
	AvailableSeatCount int                `json:"available_seat_count"`
	Fare               float64            `json:"fare"`
}

type SearchTrainRequest struct {
	DepartureStation string `json:"departure_station"`
	ArrivalStation   string `json:"arrival_station"`
	DepartureDate    string `json:"departure_date"`
}

type SearchTrainResponse struct {
	TrainNumber        string  `json:"train_number"`
	TrainName          string  `json:"train_name"`
	AvailableSeatCount int     `json:"available_seat_count"`
	Fare               float64 `json:"fare"`
}
