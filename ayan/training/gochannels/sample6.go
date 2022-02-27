package main

import "fmt"

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received Job :", j)
			} else {
				fmt.Println("Received Job :", nil)
				done <- true
				// close(jobs)
				return
			}
		}
	}()
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("Sent Job :", j)
	}
	close(jobs)
	fmt.Println("Sent Job :", nil)
	<-done
}
