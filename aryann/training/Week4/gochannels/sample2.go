package main

import "fmt"

func main() {

	msgs := make(chan string, 2)
	msgs <- "hello"
	msgs <- "world"

	fmt.Println(<-msgs)
	fmt.Println(<-msgs)

}
