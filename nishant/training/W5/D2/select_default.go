package main

import "fmt"

func main() {

	message := make(chan string, 1)

	select {
	case msg := <-message:
		fmt.Println("Msg received ", msg)
	default:
		fmt.Println("No msg received")
	}

	select {
	case message <- "hello":
		fmt.Println("Sent msg ")
	default:
		fmt.Println("Msg not sent")
	}

	select {
	case msg := <-message:
		fmt.Println("Msg received ", msg)
	default:
		fmt.Println("No msg received")
	}

}
