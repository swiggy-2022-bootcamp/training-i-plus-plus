package main

import "fmt"

func main() {

	var creature string = "shark"
	var creaturePointer *string = &creature

	fmt.Println("The value of creature is:", creature)
	fmt.Println("The address of creature is:", creaturePointer)
}
