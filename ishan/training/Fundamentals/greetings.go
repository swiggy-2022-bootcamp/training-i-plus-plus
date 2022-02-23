package main

import (
	"fmt"
)

var points = []int{20, 90, 100, 45, 70}

func sayHello(n int) {
	fmt.Println(n)
}

func newAdd() {
	var a = add(40, 35)
	fmt.Println(a)
}
