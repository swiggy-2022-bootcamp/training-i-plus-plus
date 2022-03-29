package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Doctor struct{
	Id					primitive.ObjectID		`bson:"_id,omitempty"`
	User										`json:"user" validate:"required"`
	Category 			string					`json:"category" validate:"required"`
	Yoe 	 			float64					`json:"yoe" validate:"required"`
	MedicalLicenseLink	string					`json:"medicalLicenseLink" validate:"required"`
	Appointments		[]Appointment			`json:"appointments"`	
}