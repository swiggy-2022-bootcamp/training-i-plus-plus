package models

type User struct{
	Name 		string	`json:"name" validate:"required"`
	EmailId 	string  `json:"emailId" validate:"required"`
	Age 		int		`json:"age" validate:"required"`
	Address		string  `json:"address" validate:"required"`
}