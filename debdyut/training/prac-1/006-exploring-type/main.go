package main

import "fmt"

var y = 42

var z string = "Lorem ipsum"

// raw string enclosed in ``
var a = `James said
"Shaken,

not stirred"`

// Go is a STATIC programming language
// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
// not a DYNAMIC programming language

func main() {
	fmt.Println("y:", y)
	fmt.Printf("%T\n", y)
	fmt.Println("z:", z)
	fmt.Printf("%T\n", z)

	// z holds a string value
	// below assignment to int value will throw error
	// z = 99

	fmt.Println("a:", a)
	fmt.Printf("%T\n", a)
}
