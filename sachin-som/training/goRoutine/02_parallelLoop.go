package main

import (
	"fmt"
	"time"
)

var (
	LOOP = 5
)

func main() {
	done := make(chan int, LOOP)
	for i := 0; i < LOOP; i++ {
		go func(i int) {
			fmt.Println(i, time.Now())
			done <- 1
		}(i)
	}
	for j := 0; j < LOOP; j++ {
		<-done
	}
}
