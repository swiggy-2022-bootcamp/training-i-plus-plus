package main

import (
	"fmt"
)

type human struct {
	name    string
	age     int
	address string
}

func main() {
	human1 := human{name: "ravi", age: 22, address: "new"}
	fmt.Println(human1)
	alterProperties(&human1)
	fmt.Println(human1)

	//channels
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Print(<-ch)
	ch <- 3
	fmt.Print(<-ch, <-ch)
}

func alterProperties(human1 *human) {
	(*human1).address = "changed"
}
