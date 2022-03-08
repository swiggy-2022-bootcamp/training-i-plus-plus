package main

import "fmt"

func main() {
	var name string = "Rishabh Mishra"
	var age int = 22
	var isLoggedIn bool = true

	if isLoggedIn {
		fmt.Printf("%s has age of %d has logged in\n", name, age)
	}
}
