/*
Consider the beginning of a package person in person2.go: the type Person is clearly exported, but
its fields are not! For example the statement p.firstname in use_person2.go is an error. How can we
change, or even read the name of a Person object in another program ?
*/

package main

import (
	"fmt"
	"/github.com/sachinsom93/struct"
)

type Person3 struct {
	name string
}

// Getter
func (p *Person3) FullName() (name string) {
	return p.name
}

// Setter
func (p *Person3) SetFullName(name string) {
	p.name = name
}

func main() {
	var p Person3
	var p struct.Person
	p.name = "sachin som"
	p.SetFullName("sachin som")
	fmt.Println(p.FullName())
}
