package main

import (
	"fmt"
)

func main() {
	message := make(chan string)
	signals := make(chan bool)
	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")

	}
	msg := "hi"
	select {
	case message <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-message:
		fmt.Println("Recievd mes ", msg)
	case sig := <-signals:
		fmt.Println("Recievd sig ", sig)
	default:
		fmt.Println("no activity")
	}

}
