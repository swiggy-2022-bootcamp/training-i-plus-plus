package domain

import (
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

type User struct {
	id        int
	firstName string
	lastName  string
	username  string
	password  string
	phone     string
	email     string
	role      Role
}

// ----- getters and setters --------

func (u User) Id() int {
	return u.id
}

func (u *User) SetId(id int) {
	u.id = id
}

func (u User) Email() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u User) FirstName() string {
	return u.firstName
}

func (u *User) SetFirstName(firstName string) {
	u.firstName = firstName
}

func (u User) LastName() string {
	return u.lastName
}

func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}

func (u User) Username() string {
	return u.username
}

func (u *User) SetUsername(username string) {
	u.username = username
}

func (u User) Password() string {
	return u.password
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u User) Phone() string {
	return u.phone
}

func (u *User) SetPhone(phone string) {
	u.phone = phone
}

func (u User) Role() Role {
	return u.role
}

func (u *User) SetRole(r Role) {
	u.role = r
}

func NewUser(firstName, lastName, username, phone, email, password string, role Role) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
		username:  username,
		phone:     phone,
		email:     email,
		password:  password,
		role:      role,
	}
}

type UserRepository interface {
	GetAllUsers() ([]*User, error)
	FindByUsername(string) (*User, error)
	FindByEmail(string) (*User, error)
	Save(User) (User, error)
	DeleteUserByUsername(string) (bool, error)
}
