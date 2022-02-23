package main

import "fmt"


func main () {
	sum :=0 // short declaration operator
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum is :",sum)
}