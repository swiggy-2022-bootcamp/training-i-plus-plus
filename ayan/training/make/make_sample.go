package main

import "fmt"

func main() {
	var nums = make([]int, 4, 7)
	// append()
	fmt.Println("list =", nums, "\nlength =", len(nums), "\ncapacity =", cap(nums))
}
