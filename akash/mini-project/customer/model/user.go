package model

type User struct {
	Username  string `json:"username"`
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Address   string `json:"address"`
	Telephone string `json:"tel"`
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
