package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Train struct {
	TrainId        primitive.ObjectID `json:"id,omitempty"`
	Name           string             `json:"name,omitempty" validate:"required"`
	Source         string             `json:"source,omitempty" validate:"required"`
	Destination    string             `json:"destination,omitempty" validate:"required"`
	ArrivalTime    string             `json:"arrivalTime,omitempty" validate:"required"`
	DepartureTime  string             `json:"departureTime,omitempty" validate:"required"`
	AvailableSeats []int              `json:"availableSeats,omitempty"`
	ReservedSeats  []int              `json:"reservedSeats,omitempty"`
}
