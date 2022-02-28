package main

import (
	"fmt"
	"time"
)

func display(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg, " : ", i)
	}
}

func main() {
	display("hello")

	go display("Goroutine") //go routine - light weight thread execution

	go func(name string) {
		fmt.Println("my name is", name)
	}("John")

	time.Sleep(time.Second * 10)
	fmt.Println("end")
}
