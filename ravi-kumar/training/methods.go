package main

import "fmt"

type Engineer struct {
	name string
	age  int
}

func (engineer *Engineer) UpdateName() {
	engineer.name = "Changed Name"
}

func main() {
	engineer := Engineer{name: "Ravi", age: 13}
	engineer.UpdateName()
	fmt.Println(engineer)
}
