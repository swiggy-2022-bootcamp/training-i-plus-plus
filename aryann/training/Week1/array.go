package main

import "fmt"

func main() {

	var a [2]string

	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	nums := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(nums)

	var s []int = nums[1:4]
	fmt.Println(s)
}