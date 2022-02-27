package main

import (
	"fmt"
)

type Flyer interface {
	fly() string
}

type Walker interface {
	walk() string
}

type Bird struct {
	Name string
}

func (b *Bird) fly() string {
	return "Flying..."
}

func (b *Bird) walk() string {
	return "Walking..."
}

func main() {
	var b = Bird{"Chirper"}
	// %v for value and %T for type
	fmt.Printf("%v --> %T", b, b) // {Chirper} --> main.Bird
}
