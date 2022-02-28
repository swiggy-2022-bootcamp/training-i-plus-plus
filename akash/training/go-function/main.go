package main

import "fmt"

func main() {
	fmt.Println("Hey")
	fmt.Println(adder(10, 49))
	fmt.Println(adderPro(1, 2, 3, 4, 5))
}

func adder(val1, val2 int) int {
	return val1 + val2
}

func adderPro(values ...int) (int, string) {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum, "Hello from adderPro"
}
