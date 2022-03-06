package main

import "fmt"

type User struct {
	name       string
	email      string
	age        int
	isLoggedIn bool
}

func newUser(name string, email string, age int) *User {
	user := User{name: name, email: email, age: age}
	user.isLoggedIn = true
	return &user
}

func main() {
	var user1ref = newUser("Rishabh Mishra", "rishabhmishra3321@gmail.com", 22)
	var user1 = *user1ref
	fmt.Println("user pointer", user1ref)
	fmt.Println("user", user1)
}
