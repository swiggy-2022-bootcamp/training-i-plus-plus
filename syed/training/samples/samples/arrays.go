package main

import "fmt"
var numbersArray [5]int
var namesArray [4]string

func main(){
	numbersArray[0] = 10
	numbersArray[1] = 20
	numbersArray[2] = 30
	numbersArray[3] = 40
	numbersArray[4] = 50
	fmt.Println("Array is : " ,numbersArray)
	namesArray[0] ="John"
	namesArray[1]="Jimmy"
	namesArray[2]="jackson"
	namesArray[3]="Joe"
	fmt.Println(namesArray)
}