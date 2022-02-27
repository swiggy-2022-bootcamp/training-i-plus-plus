package main

import (
	"fmt"
)
func main() {
	msgs := make(chan string,2)
	msgs <- "hello"
	msgs <- "bye"

	fmt.Println(<-msgs)
	fmt.Println(<-msgs)

	msgs <- "hello1"
	msgs <- "bye1"

	fmt.Println(<-msgs)
	fmt.Println(<-msgs)
}