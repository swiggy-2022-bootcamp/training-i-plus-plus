package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Quota int

const (
	General Quota = iota + 1
	Sleeper
	AC
)

type TicketStatus int

const (
	Reserved Quota = iota + 1
	Cancelled
	Unreserved
)

type Ticket struct {
	TicketPnr     primitive.ObjectID `json:"id,omitempty"`
	PassengerName []string           `json:"passengerName,omitempty"`
	Source        string             `json:"source,omitempty"`
	Destination   string             `json:"destination,omitempty"`
	Amount        int                `json:"amount,omitempty"`
	NumberOfSeats int                `json:"number_of_seats,omitempty"`
	TrainNumber   int                `json:"trainId,omitempty"`
	TicketStatus  TicketStatus       `json:"ticketStatus,omitempty"`
}

type TicketRequest struct {
	PassengerName []string `json:"passengerName,omitempty"`
	Source        string   `json:"source,omitempty"`
	Destination   string   `json:"destination,omitempty"`
	TrainNumber   int      `json:"trainId,omitempty"`
}
