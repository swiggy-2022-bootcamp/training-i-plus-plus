package main

import "fmt"

func main() {
	message := make(chan string)
	signal := make(chan bool)

	select {
	case msg := <-message:
		fmt.Println("receive msg ", msg)
	default:
		fmt.Println("no receive msg ")
	}

	msg := "hi"
	// message change don't have buffe that's why default statement got printed
	select {
	case message <- msg:
		fmt.Println("receive msg ", msg)
	default:
		fmt.Println("no receive msg ")
	}

	select {
	case msg := <-message:
		fmt.Println("receive msg ", msg)
	case msg := <-signal:
		fmt.Println("receive sig ", msg)
	default:
		fmt.Println("no activity ")
	}

}
