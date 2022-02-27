package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status string
	Age    string
}

func (user *User) updateEmail(newEmail string) {
	fmt.Println("old email: ", user.Email)
	user.Email = newEmail
	fmt.Println("new email: ", user.Email)
}

func main() {
	murtaza := User{"Murtaza", "murtaza@hi.com", "online", "24"}
	murtaza.updateEmail("murtaza@hello.com")

	fmt.Println(murtaza)
}
