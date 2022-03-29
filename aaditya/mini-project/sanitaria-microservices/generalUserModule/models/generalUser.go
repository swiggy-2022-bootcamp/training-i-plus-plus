package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GeneralUser struct{
	Id 					primitive.ObjectID		`bson:"_id,omitempty"`
	User										`json:"user" validate:"required"`
	PreviousDiseases	string					`json:"previousDisease" validate:"required"`
	IsPatient			bool					`json:"isPatient"`	
	Appointments		[]Appointment			`json:"appointments"`
}

