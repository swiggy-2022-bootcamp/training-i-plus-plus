package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	p1 := person{
		firstName: "James",
		lastName:  "Bond",
		age:       45,
	}

	fmt.Println(p1)

}
