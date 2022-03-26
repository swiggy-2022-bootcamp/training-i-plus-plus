package main

import (
	"fmt"
)

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	SimpleOperation()
	UnblockedOperation()
}

func SimpleOperation() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "Hello")
	pong(pings, pongs)
	fmt.Print(<-pongs)
}

func UnblockedOperation() {
	message := make(chan string)
	signal := make(chan bool)
	select {
	case msg := <-message:
		fmt.Println("Received message", msg)
	default:
		fmt.Println("No message received")
	}
	msg := "hi"
	select {
	case message <- msg:
		fmt.Println("Received message", msg)
	default:
		fmt.Println("No message received")
	}

	select {
	case msg := <-message:
		fmt.Println("Received message", msg)
	case msg := <-signal:
		fmt.Println("Received signal", msg)
	default:
		fmt.Println("Nothing received")
	}
}
