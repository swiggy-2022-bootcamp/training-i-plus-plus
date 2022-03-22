package main

import "fmt"

var flag1, flag2, flag3 bool

func add(a int, b int) int {
	return a + b
}

func main() {

	fmt.Println("Sum: ", add(10, 20))

	var x int
	fmt.Println("Variable values: ", x, flag1, flag2, flag3)
}
