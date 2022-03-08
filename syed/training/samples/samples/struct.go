package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	//Assigning values to the fields in the person struct:
	p1 := person{
		firstName: "Mark",
		lastName:  "Kedu",
		age:       30,
	}

	fmt.Println("The is the person: ", p1)
}