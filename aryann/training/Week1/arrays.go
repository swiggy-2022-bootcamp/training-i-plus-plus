package main

import "fmt"

var numbersArray [5]int
var namesArray [5]string

func main() {

	numbersArray[0] = 10
	numbersArray[1] = 20
	numbersArray[2] = 30
	numbersArray[3] = 40
	numbersArray[4] = 50

	fmt.Println("Number Array:", numbersArray)

	namesArray[0] = "John1"
	namesArray[1] = "John2"
	namesArray[2] = "John3"
	namesArray[3] = "John4"
	namesArray[4] = "John5"

	fmt.Println("Names Array:", namesArray)
}
