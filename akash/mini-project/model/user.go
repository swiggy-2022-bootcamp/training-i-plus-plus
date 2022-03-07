package model

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type LoginData struct {
	Email    string
	Password string
}

type SignupData struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}
