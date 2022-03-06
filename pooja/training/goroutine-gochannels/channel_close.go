package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("recevied job", j)
			} else {
				done <- true
			}
		}
	}()
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	if <-done {
		fmt.Println("received all jobs")
	}
}
