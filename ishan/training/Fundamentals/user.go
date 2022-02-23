package main

import "fmt"

type user struct {
	name     string
	email    string
	password string
	address  string
}

var users = map[string]user{}

func createNewUser(name string, email string, password string, address string) user {
	u := user{
		name:     name,
		address:  address,
		password: password,
		email:    email,
	}
	users[email] = u
	return u
}

func (u user) getUserInfo() string {

	return fmt.Sprintf(" Name : %v \n Email : %v \n Password : %v \n Address : %v\n", u.name, u.email, u.password, u.address)
}

func logInUser(email string, password string) bool {
	if users[email].password == password {
		return true
	}
	return false
}
