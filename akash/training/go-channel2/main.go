package main

import "fmt"

func main() {
	msgs := make(chan string, 2)
	msgs <- "a"
	msgs <- "b"
	fmt.Println(<-msgs)
	msgs <- "c"
	fmt.Println(msgs)
	fmt.Println(<-msgs)
}
