package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	sendData(ch)
	receiveData(ch)
	time.Sleep(1e9)
	fmt.Println("End of main function.")
}

func sendData(ch chan string) {
	ch <- "sachin"
	ch <- "som"
	ch <- "jaipur"
	ch <- "rajasthan"
}

func receiveData(ch chan string) {

}
