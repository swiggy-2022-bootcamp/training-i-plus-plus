package main

import (
	"fmt"
)

type User struct {
	username string
	password string
}

func (u User) user_details() {
	fmt.Printf("Username is %s\n", u.username)
	fmt.Printf("Password is %s\n", u.password)
	fmt.Println("===========================")
}

func main() {
	u1 := User{"admin", "admin123"}
	u1.user_details()

	u2 := User{"John", "123456"}
	u2.user_details()
}
