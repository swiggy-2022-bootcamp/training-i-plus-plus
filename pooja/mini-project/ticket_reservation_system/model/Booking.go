package model

import "go.mongodb.org/mongo-driver/bson/primitive"

var GlobalPNR = int64(1)

type BookTicketsRequest struct {
	TrainNumber   string `json:"train_number"`
	DepartureDate string `json:"departure_date"`
	NumberOfSeats int    `json:"seat_count"`
	UserName      string `json:"username"`
}

type Booking struct {
	ID       primitive.ObjectID `json:"id" bson:"id"`
	UserName string             `json:"username"`
	//add list of seats
	PNR           int64  `json:"pnr"`
	NumberOfSeats int    `json:"seat_count"`
	TrainNumber   string `json:"train_number"`
	DepartureDate string `json:"departure_date"`
	BookingStatus string `json:"booking_status"`
}

type CancelBookingRequest struct {
	PNR      int64  `json:"pnr"`
	UserName string `json:"username"`
	//add pnr number
}

type BookingsByUserRequest struct {
	UserName string `json:"username"`
}
