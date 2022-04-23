package main

import "fmt"

//User ..
type User struct {
	Name  string
	Age   string
	Email string
}

func (user *User) modifyFunc(name string) {
	fmt.Println("Existing Name: ", user.Name)
	user.Email = name
	fmt.Println("new Name: ", user.Name)
}

func main() {
	ashwin := User{"Ashwin Gopalsamy", "22", "ashwin@swiggy.com"}
	ashwin.modifyFunc("AG")

	fmt.Println(ashwin)
}
