package main

import (
    "fmt"
)

func main() {
	sliceOne := []int {0,1,2,3,4,5}
	fmt.Println(sliceOne)

	sliceTwo := sliceOne[2:4]
	fmt.Println(sliceTwo)

	fmt.Println("length: ", len(sliceTwo))
	fmt.Println("capacity: ", cap(sliceTwo))

}
