package main

import "fmt"

func main() {

	var nums = make([]int, 4, 7)
	fmt.Printf("numbers: %v, \n length = %d, \n capacity = %d", nums, len(nums), cap(nums))
}
