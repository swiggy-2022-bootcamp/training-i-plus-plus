package main

import (
    "strconv"
    "fmt"
)

func main() {
	//declaration
	var firstArray [5]string
	fmt.Println("first array: ",firstArray) 
	//assignment
	firstArray[0] = "Hello"
	fmt.Println("first array: ",firstArray) 

	for i := 1; i<5 ; i++{
		firstArray[i] = strconv.Itoa(i*5)
	}
	fmt.Println("first array: ",firstArray)

	//initialization
	var secondArray = [3]int{1,2,3}
	fmt.Println("second array: ",secondArray)

	//short hand declaration
	thirdArray := [4]int{6,7,8,9}
	fmt.Println("third array: ",thirdArray)
}
