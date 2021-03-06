package models

type SignUp struct {
	Username string `json:"username,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required"`
	TypeOf   string `json:"typeOf,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}
