package model

type Role string

const (
	SELLER Role = "SELLER"
	BUYER  Role = "BUYER"
)

type User struct {
	Email     string `json:"email" bson:"email" example:"sd@gmail.com"`
	Firstname string `json:"firstname" bson:"firstname" example:"Sujata"`
	Lastname  string `json:"lastname" bson:"lastname" example:"Dwivedi"`
	Password  string `json:"password" bson:"password" example:"password"`
	Address   string `json:"address" bson:"address" example:"India"`
	// Enum for user role type (BUYER or SELLER)
	Role Role `json:"role" bson:"role" example:"SELLER"`
}

type SigninUser struct {
	Email    string `json:"email" bson:"email" example:"sd@gmail.com"`
	Password string `json:"password" bson:"password" example:"password"`
}
