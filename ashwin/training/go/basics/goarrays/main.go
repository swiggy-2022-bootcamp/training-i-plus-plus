package main

import (
	"fmt"
)

func array1() {

	// Declare Array using make
	goarr := make([]string, 10)

	goarr[0] = "apple"
	goarr[1] = "orange"
	goarr[2] = "grapes"

	fmt.Println("Content", goarr, "Length", len(goarr))

	// Declare and Initialize array
	arr2 := [5]float32{10.1, 34.6, 77.9, 80.3, 11.1}
	fmt.Println("Content", arr2, "Length", len(arr2))
}

func main() {
	array1()
}
