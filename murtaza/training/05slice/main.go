package main

import "fmt"

/**
 * Slices are most widely used Data structure in Go
 * Internally makes use of array
 * dynamic memory allocation (reallocation)
 *
 */

func DeclareSlice() {
	myslice := []string{}
	myslice = append(myslice, "carrot")
	myslice = append(myslice, "cucumber")
	myslice = append(myslice, "tomato")

	fmt.Println(myslice)
}

func RemoveElementFromSliceAtIndex(index int) {
	var myslice = []string{"apple", "mango", "banana"}
	var fruitList = []string{}

	// initialize a slice using other slice
	fruitList = append(fruitList, myslice...)
	fmt.Println(fruitList)

	//remove element from slice at index
	fmt.Println(fruitList[:index], fruitList[index+1:])
	fruitList = append(fruitList[:index], fruitList[index+1:]...)
	fmt.Println(fruitList)
}

func main() {
	DeclareSlice()
	RemoveElementFromSliceAtIndex(1)
}
