package main

import (
	"fmt"
	"time"
)

func worker(id int, job <-chan int, result chan<- int) {

	for j := range job {
		fmt.Println("Worker ", id, " started job ", j)

		time.Sleep(500 * time.Millisecond)

		fmt.Println("Worker ", id, " completed job ", j)
		result <- j * 10
	}
}

func main() {

	numWork := 3
	numJobs := 10
	jobs := make(chan int)
	result := make(chan int, 10)

	for i := 1; i <= numWork; i++ {
		go worker(i, jobs, result)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}

	for i := 1; i <= numJobs; i++ {
		fmt.Println("result ", <-result)
	}
}
