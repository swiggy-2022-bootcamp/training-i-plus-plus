package main

import (
	"fmt"
	"time"
)

func display(msg string) {

	for i := 0; i < 10; i++ {
		fmt.Println(msg, ":", i)
	}
}

func main() {

	display("Hello")

	go display("World")

	go func(name string) {
		fmt.Println("My name is:", name)
	}("Golang")

	time.Sleep(time.Second)
	fmt.Println("Done")

}
