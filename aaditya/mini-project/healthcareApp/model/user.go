package model

type User struct{
	Id			string	`json:"id"`
	Name 		string	`json:"name"`
	EmailId 	string  `json:"emailId"`
	Age 		int		`json:"age"`
	Address		string  `json:"address"`
}