package main

import "fmt"

type Example struct {
	val int
}

func (e *Example) changeValue(newVal int) {
	e.val = newVal
}

func (e Example) printValue() {
	fmt.Println("Inside call by value printValue.")
	fmt.Println(e.val)
}

// NOT OK: method is already defined for Example struct
// func (e *Example) printValue() {
// 	fmt.Println("Inside call by reference printValue.")
// }

func main() {
	var e Example     // Struct Value
	e.changeValue(12) // Go itself pass the reference to the method
	e.printValue()
}
