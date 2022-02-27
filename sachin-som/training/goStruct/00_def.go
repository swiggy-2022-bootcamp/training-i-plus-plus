package main

import "fmt"

/**
 * Structs in go represt a real world entity.
 * It is a composite type with some properties.
 * type indetifier struct {
	 field1 type1
	 field2 type2
	 ...
 }

 * Shorthand
 * type indentifier struct ( a, b int )
*/

type Person struct {
	name string
	age  uint
}

func main() {
	fmt.Println("Working...")
}
