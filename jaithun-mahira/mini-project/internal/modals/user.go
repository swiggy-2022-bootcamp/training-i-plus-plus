package modals

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Contact ContactInfo `json:"contact"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	IdProof string `json:"idProof"`
	Role Role `json:"role"`
}

type ContactInfo struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Role int

const (
    Admin Role = iota 
    Customer 
)