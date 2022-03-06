package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJob = 5
	jobs := make(chan int, numJob)
	results := make(chan int, numJob)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= numJob; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= numJob; a++ {
		<-results
	}
}
