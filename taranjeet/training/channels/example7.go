package main

import (
	"fmt"
	"time"
)

func test() {
	fmt.Println("middle")
}

//worker 3 started  job 1
//worker 1 started  job 2
//worker 2 started  job 3
//worker 2 finished job 3
//worker 3 finished job 1
//worker 3 started  job 5
//worker 2 started  job 4
//worker 1 finished job 2
//worker 3 finished job 5
//worker 2 finished job 4

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		//fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		//fmt.Println("worker", id, "finished job", j)
		results <- j * 2
		fmt.Println("sent ")

	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		x := <-results
		fmt.Printf("received value - %d\n", x)
	}

	jobs2 := make(chan int, 5)

	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs2
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all job")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs2 <- j
		fmt.Println("sent job", j)
	}

	close(jobs2)
	fmt.Println("sent all job")
	<-done

}
