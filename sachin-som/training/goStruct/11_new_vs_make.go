package main

import "fmt"

// Map
type Foo map[string]string

// Struct
type Bar struct {
	thingOne string
	thingTwo int
}

func main() {

	// New can be used for structs
	b := new(Bar)
	(*b).thingOne = "sachin"
	(*b).thingTwo = 2

	// Make can't be used for structs
	// b2 := make(Bar) // Compile time error

	// New can't be used for maps
	// m := new(Foo)
	// (*m)["x"] = "s"
	// (*m)["y"] = "ss"
	// fmt.Println(m) // Runtime Error => panic: assignment to entry in nil map

	// Make can be used for maps
	m := make(Foo)
	m["x"] = "sachin"
	m["y"] = "som"
	fmt.Println(m)
}
