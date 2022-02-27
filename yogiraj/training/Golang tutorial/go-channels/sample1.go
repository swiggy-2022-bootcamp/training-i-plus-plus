package main

import "fmt"

func main(){
	msgs := make(chan string)

	go func(){
		msgs <-"hello"
	}()

	fmt.Println(<-msgs)
}