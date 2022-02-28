package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, isOpen := <-jobs
			if isOpen {
				fmt.Println("received ", j)
			} else {
				fmt.Println("channel closed")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i * 10
		fmt.Println("Sent ", i*10)
	}

	close(jobs)
	<-done
}
