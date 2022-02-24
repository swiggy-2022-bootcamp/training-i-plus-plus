package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	// fatal error: all goroutines are asleep - deadlock!
	// sendData(ch)
	// receiveData(ch)
	go sendData(ch)
	go receiveData(ch)
	time.Sleep(12e9)
	fmt.Println("End of main function.")
}

// Send will wait 2s after each insertion
// untill receiveData receives the string
func sendData(ch chan string) {
	ch <- "sachin"
	ch <- "som"
	ch <- "jaipur"
	ch <- "rajasthan"
}

func receiveData(ch chan string) {
	var msg string
	for {
		msg = <-ch
		time.Sleep(2e9)
		fmt.Println(msg)
	}
}
