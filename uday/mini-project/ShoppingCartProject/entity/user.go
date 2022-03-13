package entity

type User struct{
	Id int   `json:"id"`
	Username string `json:"username"`
	Password string  `json:"password"`
	Email string	`json:"email"`
	Phone int		`phone:"phone"`
}

