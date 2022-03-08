package main

import "fmt"

func main() {
	mySlice := make([]int, 4, 10)
	fmt.Printf("%v, %d, %d \n", mySlice, len(mySlice), cap(mySlice))

	mySlice = append(mySlice, 10, 20)

	for k, v := range mySlice {
		fmt.Println(k, v)
	}
}
