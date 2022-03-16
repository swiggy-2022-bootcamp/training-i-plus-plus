package model

import (
	"github.com/google/uuid"
)

const (
	UserTableName = "users"
)

type User struct {
	userId   string
	email    string
	password string
}

func NewUser(email, pass string) *User {
	return &User{
		userId:   uuid.New().String(),
		email:    email,
		password: pass,
	}
}

func (u *User) GeneraterId() {
	u.userId = uuid.New().String()
}
