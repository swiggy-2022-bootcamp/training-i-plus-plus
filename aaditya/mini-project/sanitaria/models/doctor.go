package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Doctor struct{
	Id		primitive.ObjectID		`json:"id,omitempty"`
	User							`json:"user"`
	Category 			string		`json:"category"`
	Yoe 	 			float64		`json:"yoe"`
	MedicalLicenseLink	string		`json:"medicalLicenseLink"`
}