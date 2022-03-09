package main

import "fmt"

func returnNames() (string, string) {
	return "Rishabh", "Mishra"
}

func main() {
	firstName, lastName := returnNames()
	fmt.Println("Hi", firstName, lastName)
}
