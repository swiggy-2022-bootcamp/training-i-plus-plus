package main

import "fmt"

func main() {

	sum := 0
	const n int = 5
	var arr [n]int

	i := 0
	for i < 5 {
		arr[i] = i
		sum += arr[i]
		i++
	}

	fmt.Println("Array", arr)
	fmt.Println("Sum", sum)

}
