package main

import "fmt"

var v1,v2 bool
var a = "some String"
func main() {
    fmt.Println("Nishant OP! ", p1(4, 6))
	fmt.Println(v1, v2)

	
	b := "other string"
	fmt.Println(a, "\n", b)

	n := 10
	fmt.Printf("sum of %d = %d", n, loop(n))
}

func p1(a int, b int) int {
	return a*b
}


func loop(n int) (sum int) {

	for i := 1; i<=n; i++ {
		sum += i;
	}
	return
}