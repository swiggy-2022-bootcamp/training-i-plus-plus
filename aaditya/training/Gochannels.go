package main

import (
	"fmt"
)

func main(){
	msgs := make(chan string)
	go func(){
		msgs <- "hello"
		fmt.Println("anonymous function executed")
	}()

	res:= <-msgs
	fmt.Println("Received msg from goroutine")
	fmt.Println(res)
}
