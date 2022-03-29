package domain

import (
	"encoding/json"
	"errors"
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

func GetEnumByIndex(idx int) (Role, error) {
	switch idx {
	case 0:
		return Admin, nil
	case 1:
		return Seller, nil
	case 2:
		return Customer, nil
	default:
		return -1, errors.New("invalid enum index")
	}
}

// type User struct {
// 	id        int
// 	firstName string
// 	lastName  string
// 	username  string
// 	password  string
// 	phone     string
// 	email     string
// 	role      Role
// }
type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Role      Role   `json:"role"`
}

// ----- getters and setters --------

// func (u User) Id() int {
// 	return u.id
// }

func (u *User) SetId(id int) {
	u.Id = id
}

// func (u User) Email() string {
// 	return u.email
// }

func (u *User) SetEmail(email string) {
	u.Email = email
}

// func (u User) FirstName() string {
// 	return u.firstName
// }

func (u *User) SetFirstName(firstName string) {
	u.FirstName = firstName
}

// func (u User) LastName() string {
// 	return u.lastName
// }

func (u *User) SetLastName(lastName string) {
	u.LastName = lastName
}

// func (u User) Username() string {
// 	return u.username
// }

func (u *User) SetUsername(username string) {
	u.Username = username
}

// func (u User) Password() string {
// 	return u.password
// }

func (u *User) SetPassword(password string) {
	u.Password = password
}

// func (u User) Phone() string {
// 	return u.phone
// }

func (u *User) SetPhone(phone string) {
	u.Phone = phone
}

// func (u User) Role() Role {
// 	return u.role
// }

func (u *User) SetRole(r Role) {
	u.Role = r
}

func (u User) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":        u.Id,
		"firstName": u.FirstName,
		"lastName":  u.LastName,
		"email":     u.Email,
		"password":  u.Password,
		"username":  u.Username,
		"phone":     u.Phone,
		"role":      u.Role,
	})
}

func NewUser(firstName, lastName, username, phone, email, password string, role Role) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Phone:     phone,
		Email:     email,
		Password:  password,
		Role:      role,
	}
}

type UserRepository interface {
	GetAllUsers() ([]*User, error)
	FindByUsername(string) (*User, error)
	FindByEmail(string) (*User, error)
	FindByUserId(int) (*User, error)
	Save(User) (User, error)
	DeleteUserByUsername(string) (bool, error)
}

type UserMongoRepository interface {
	InsertUser(User) (User, error)
	FindUserById(int) (*User, error)
	FindUserByUsername(string) (*User, error)
	DeleteUserByUserId(int) error
	UpdateUser(User) (*User, error)
}
