package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Train struct {
    Id             primitive.ObjectID `json:"id,omitempty"`
    Name           string             `json:"name,omitempty" validate:"required"`
    Source         string             `json:"source,omitempty" validate:"required"`
    Destination    string             `json:"destination,omitempty" validate:"required"`
		TotalSeats     int64              `json:"totalseats,omitempty" validate:"required"`
		AvailableSeats int64              `json:"availableseats,omitempty" validate:"required"`
    ArrivalTime    time.Time          `json:"arrivaltime,omitempty" validate:"required"`
    DepartureTime  time.Time          `json:"departuretime,omitempty" validate:"required"`
    TicketPrice    float32            `json:"ticketprice,omitempty" validate:"required"`
}