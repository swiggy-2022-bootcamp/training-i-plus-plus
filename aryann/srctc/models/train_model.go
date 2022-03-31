package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Train struct {
	Id          primitive.ObjectID `json:"id,omitempty" validate:"required"`
	Source      string             `json:"source,omitempty" validate:"required"`
	Destination string             `json:"destination,omitempty" validate:"required"`
}
