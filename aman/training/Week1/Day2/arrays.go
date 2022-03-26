package main

import "fmt"

var numbersArray [5]int
var namesArray [4]string

func main() {
	numbersArray[0] = 10
	numbersArray[1] = 20
	numbersArray[2] = 30
	numbersArray[3] = 40
	numbersArray[4] = 50
	fmt.Println("Array is : ", numbersArray)

	namesArray[0] = "John"
	namesArray[1] = "Sam"
	namesArray[2] = "Joe"
	fmt.Println("Array is : ", namesArray)

	nums := [5]int{2, 3, 4, 5, 6}

	fmt.Println(nums)

}
