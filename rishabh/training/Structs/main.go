package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipcode int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateFirstName(newFirstname string) {
	p.firstname = newFirstname
}

func main() {
	jim := person{
		firstname: "Jim",
		lastname:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipcode: 201012,
		},
	}
	jim.updateFirstName("Rishabh")
	jim.print()
}
