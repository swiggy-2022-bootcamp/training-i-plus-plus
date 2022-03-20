package models

type Address struct{
	State 	string	`json:"state" 	bson:"state"`
	City 	string	`json:"city" 	bson:"city"`
	Pincode	int		`json:"pincode" bson:"pincode"`
}

type User struct{
	Name 	string 	`json:"name" 	bson:"user_name"`
	Age 	int		`json:"age" 	bson:"user_age"`
	Address Address	`json:"äddress"	bson:"user_address"`
}