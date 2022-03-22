package main

import (
	"fmt"
)

type User struct {
	username string
	password string
}

func (user User) user_details() {

	fmt.Println("Username: ", user.username)
	fmt.Println("Password: ", user.password)
}

func main() {

	u1 := User{"aryann", "somepwd"}
	u1.user_details()

	u2 := User{"dhir", "somepwd2"}
	u2.user_details()
}
