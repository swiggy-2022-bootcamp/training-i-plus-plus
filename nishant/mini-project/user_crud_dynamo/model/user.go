package model

import (
	"github.com/google/uuid"
)

const (
	UserTableName = "users"
)

var defaultUserProjection string = "userId, e	mail"

type User struct {
	UserId   string `json:"userId,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func NewUser(email, pass string) *User {
	return &User{
		UserId:   uuid.New().String(),
		Email:    email,
		Password: pass,
	}
}

func (u *User) GeneraterId() {
	u.UserId = uuid.New().String()
}

func GetDefaultUserProjection() *string {
	return &defaultUserProjection
}
