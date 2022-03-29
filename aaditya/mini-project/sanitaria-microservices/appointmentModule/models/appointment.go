package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Appointment struct {
	Id				primitive.ObjectID	`bson:"_id,omitempty"`
	Slot			string				`json:slot`
	Fees			int					`json:fees`
	Occupied		bool				`json:occupied`
	DoctorID		primitive.ObjectID	`json:doctorId`
	GeneralUserID	primitive.ObjectID	`json:generalUserId`
}