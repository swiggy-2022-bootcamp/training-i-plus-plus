package domain

import (
	"encoding/json"
	"panem/utils/errs"
)

type Role int

const (
	Admin    Role = iota // EnumIndex = 0
	Seller               // EnumIndex = 1
	Customer             // EnumIndex = 2
)

func (r Role) String() string {
	return [...]string{"admin", "seller", "customer"}[r]
}

func (r Role) EnumIndex() int {
	return int(r)
}

func GetEnumByIndex(idx int) (Role, *errs.AppError) {
	switch idx {
	case 0:
		return Admin, nil
	case 1:
		return Seller, nil
	case 2:
		return Customer, nil
	default:
		return -1, errs.NewUnexpectedError("invalid enum index")
	}
}

type User struct {
	Id              int    `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	Role            Role   `json:"role"`
	PurchaseHistory []int  `json:"purchase_history"`
}

func (u User) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":               u.Id,
		"firstName":        u.FirstName,
		"lastName":         u.LastName,
		"email":            u.Email,
		"password":         u.Password,
		"username":         u.Username,
		"phone":            u.Phone,
		"role":             u.Role,
		"purchase_history": u.PurchaseHistory,
	})
}

func NewUser(firstName, lastName, username, phone, email, password string, role Role) *User {
	return &User{
		FirstName:       firstName,
		LastName:        lastName,
		Username:        username,
		Phone:           phone,
		Email:           email,
		Password:        password,
		Role:            role,
		PurchaseHistory: make([]int, 0),
	}
}

type UserMongoRepository interface {
	InsertUser(User) (User, *errs.AppError)
	FindUserById(int) (*User, *errs.AppError)
	FindUserByUsername(string) (*User, *errs.AppError)
	DeleteUserByUserId(int) *errs.AppError
	UpdateUser(User) (*User, *errs.AppError)
}
