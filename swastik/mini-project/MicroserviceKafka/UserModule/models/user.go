package models

type Address struct{
	State 	string	`json:"state"`
	City 	string	`json:"city"`
	Pincode	int		`json:"pincode"`
}

type User struct{
	Name 	string 	`json:"name"`
	Age 	int		`json:"age"`
	Address Address	`json:"address"`
	Username string	`json:"username"`
	Password string	`json:"password"`
}

type Login struct{
	Username string	`json:"username"`
	Password string	`json:"password"`
}