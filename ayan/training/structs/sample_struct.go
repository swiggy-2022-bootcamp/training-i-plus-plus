package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p1 := person{
		firstName: "Ayan",
		lastName:  "Dutta",
		age:       23,
	}
	fmt.Println("Hey! This is me:", p1)
}
