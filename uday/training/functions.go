package main

import "fmt"

func fun(a int, b int) (n int, m int) {
	n = a
	m = b
	return
}
func main() {
	a, b := fun(3, 4)
	fmt.Println(a, b)
}
