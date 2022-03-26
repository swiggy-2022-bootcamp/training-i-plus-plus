package main

import "fmt"

func main() {
	intValue()
	floatValue()
	complexValue()
	byteValue()
	runeValue()
	stringValue()
	boolValue()
	mapValue()
	sliceValue()
	funcValue()
	pointerValue()
}

func intValue() {
	var a int
	var b uint
	fmt.Println("Default value of int is ", a)
	fmt.Println("Default value of uint is ", b)
}

func floatValue() {
	var a float32
	fmt.Println("Default value of float is ", a)
}

func complexValue() {
	var a complex64
	fmt.Println("Default value of complex is ", a)
}

func byteValue() {
	var a byte
	fmt.Println("Default value of byte is ", a)
}

func runeValue() {
	var a rune
	fmt.Println("Default value of rune is ", a)
}

func stringValue() {
	var a string
	fmt.Println("Default value of string is ", a)
	fmt.Println("Length of default value of string is ", len(a))
}

func boolValue() {
	var a bool
	fmt.Println("Default zero value of bool is ", a)
}

func mapValue() {
	var a map[bool]bool
	fmt.Println(a == nil)
	fmt.Println("Default value of map is ", a)
}

func sliceValue() {
	var a []int
	fmt.Println(a == nil)
	fmt.Println("Default value of slice is ", a)
}

func funcValue() {
	var a func()
	fmt.Println("Default value of a func is ", a)
}

func pointerValue() {
	var a *int
	fmt.Println("Default value of a pointer is ", a)
}
