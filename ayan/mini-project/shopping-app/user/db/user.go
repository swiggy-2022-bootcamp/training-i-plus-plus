package db

import "time"

type User struct {
	id        int
	email     string
	password  string
	name      string
	address   string
	zipcode   int32
	mobileNo  string
	role      string
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(email string, password string, name string, address string, zipcode int32, mobileNo string, role string) *User {
	return &User{
		email:    email,
		password: password,
		name:     name,
		address:  address,
		zipcode:  zipcode,
		mobileNo: mobileNo,
		role:     role,
	}
}

func (u *User) SetId(id int) {
	u.id = id
}

func (u *User) SetCreatedAt(createdAt time.Time) {
	u.createdAt = createdAt
}

func (u *User) SetUpdatedAt(updatedAt time.Time) {
	u.updatedAt = updatedAt
}
