package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ticket struct {
    Id             primitive.ObjectID `json:"id,omitempty"`
    TrainNumber    string             `json:"trainnumber,omitempty" validate:"required"`
    NoOfSeats      int32              `json:"noofseats,omitempty" validate:"required"`
    UserName       string             `json:"username,omitempty" validate:"required"`
		Passengers 		 []string           `json:"passengers,omitempty" validate:"required"`
		TotalCost      float64            `json:"totalcost,omitempty"`
    Status         string             `json:"status,omitempty"`
}


type Train struct {
	Id             primitive.ObjectID `json:"id,omitempty"`
	Name           string             `json:"name,omitempty" validate:"required"`
	Source         string             `json:"source,omitempty" validate:"required"`
	Destination    string             `json:"destination,omitempty" validate:"required"`
	TotalSeats     int32              `json:"totalseats,omitempty" validate:"required"`
	AvailableSeats int32              `json:"availableseats,omitempty" validate:"required"`
	ArrivalTime    time.Time          `json:"arrivaltime,omitempty" validate:"required"`
	DepartureTime  time.Time          `json:"departuretime,omitempty" validate:"required"`
	TicketPrice    float32            `json:"ticketprice,omitempty" validate:"required"`
}