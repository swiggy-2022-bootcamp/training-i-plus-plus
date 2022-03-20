package model

type Role int

const (
	SELLER Role = iota
	BUYER
)

type User struct {
	Email     string `json:"email" bson:"email"`
	Firstname string `json:"firstname" bson:"firstname"`
	Lastname  string `json:"lastname" bson:"lastname"`
	Password  string `json:"password" bson:"password"`
	Address   string `json:"address" bson:"address"`
	// Enum for user role type (BUYER or SELLER)
	Role Role `json:"role" bson:"role"`
}
