package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Name        string             `json:"name,omitempty" validate:"required"`
	Password    string             `json:"password,omitempty" validate:"required"`
	Contact     ContactInfo        `json:"contact,omitempty" validate:"required"`
	DateOfBirth time.Time          `json:"dateOfBirth,omitempty" validate:"required"`
	IdProof     string             `json:"idProof,omitempty" validate:"required"`
	Role        Role               `json:"role"`
}

type ContactInfo struct {
	Email string `json:"email,omitempty" validate:"required"`
	Phone string `json:"phone,omitempty" validate:"required"`
}

type Role int

const (
    Admin Role = iota 
    Customer 
)