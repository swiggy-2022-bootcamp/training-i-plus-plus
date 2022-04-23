package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Doctor Model ..
type Doctor struct {
	DoctorID                primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DoctorName              string             `json:"doctorName"`
	DoctorSpecialization    string             `json:"specialization"`
	Phone                   string             `json:"phoneNumber"`
	VerifiedDoctor          bool               `json:"verifiedDoctor"`
	ProvidesPremiumServices bool               `json:"providesPremiumServices"`
	Address                 string             `json:"address"`
	Timings                 primitive.DateTime `json:"timings" bson:"timings"`
}
