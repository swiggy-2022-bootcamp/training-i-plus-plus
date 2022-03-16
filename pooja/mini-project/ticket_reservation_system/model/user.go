package model

type User struct {
	UserId   string `json:"user_id"`
	UserName string `json:"username"`
	EmailId  string `json:"email_id"`
	Password string `json:"password"`
}
