package main

import "fmt"

func receiveMsg(c8 chan string) {
	select {
	case msg := <-c8:
		fmt.Println("message received from c8- ", msg)
	default:
		fmt.Println("no message")

	}

	//fmt.Println("message received", <-c8)
}

func main() {
	c8 := make(chan string)
	//c9 := make(chan string)

	go receiveMsg(c8)
	c8 <- "hi"
	select {
	case c8 <- "hi":
		fmt.Println("message sent to c8")
	default:
		fmt.Println("no message")
	}
	select {
	case msg := <-c8:
		fmt.Println("message received from c8- ", msg)
	default:
		fmt.Println("no message")

	}

}
