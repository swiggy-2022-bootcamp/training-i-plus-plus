package main

import "fmt"

type user struct {
	username string
	password string
}

func (u user) user_details(){
	fmt.Printf("Username is %s\n",u.username)
	fmt.Printf("Password is %s\n",u.password)
	fmt.Printf("===========================\n")
}



func main() {
	u1 := user{"admin", "admin123"}
	u1.user_details()

	u2 := user{"John", "123456"}
	u2.user_details()
}