package model

import (
	"github.com/google/uuid"
)

const (
	UserTableName = "users"
)

var defaultUserProjection string = "userId, email, username"

type User struct {
	UserId string `json:"userId,omitempty"`
	Email  string `json:"email,omitempty"`
	Pass   []byte `json:"pass,omitempty""`
	Name   string `json:"username,omitempty"`
}

func NewUser(name, email string, pass []byte) *User {
	return &User{
		UserId: uuid.New().String(),
		Email:  email,
		Pass:   pass,
		Name:   name,
	}
}

func (u *User) GeneraterId() {
	u.UserId = uuid.New().String()
}

func GetDefaultUserProjection() *string {
	return &defaultUserProjection
}
