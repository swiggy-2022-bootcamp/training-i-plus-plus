package models

type User struct {
	Name  string `json:"name,omitempty" validate:"required"`
	Email string `json:"email,omitempty" validate:"required"`
}
