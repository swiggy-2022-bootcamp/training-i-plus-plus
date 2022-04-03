package dto

type UserLoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
