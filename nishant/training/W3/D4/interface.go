package main

import "fmt"

type Animal interface {
	speak()
}

type Cat struct {
	voice string
}

type Dog struct {
	voice string
}

func (c *Cat) speak() {
	fmt.Printf("this is a cat : %s \n", c.voice)
}

func (d *Dog) speak() {
	fmt.Printf("this is a dog : %s \n", d.voice)
}

func info(a *Animal) {
	a.speak()
}

func main() {
	d := Dog{"bhow"}
	c := Cat{"meow"}

	// d.speak()
	// c.speak()

	info(d)
	info(c)
}
