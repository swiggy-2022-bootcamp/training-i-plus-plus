package main

import "fmt"

var numbers [5]int


func main(){
	numbers[0] = 1
	numbers[1] = 4
	numbers[2] = 9
	numbers[3] = 16
	numbers[4] = 25

	fmt.Println(numbers)
}