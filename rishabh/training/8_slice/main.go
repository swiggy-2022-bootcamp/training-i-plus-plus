package main

import "fmt"

func main() {
	name := make([]string, 2)
	name[0] = "Mr."
	name[1] = "Rishabh"

	name = append(name, "Mishra")

	fmt.Println("Slice is ", name)

	salutation := name[:1]
	actualName := name[1:]

	fmt.Println("Salutation", salutation)
	fmt.Println("Actual Name", actualName)

	cpyName := make([]string, len(name))
	copy(cpyName, name)

	fmt.Println("Copyied Slice is ", cpyName)
}
