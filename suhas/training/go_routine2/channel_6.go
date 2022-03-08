package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("recived job",j)
			} else {
				fmt.Println("recieved all jobs")
				done <- true
				return 
			}
		}
	}()
	for j :=1 ; j<=3 ;j ++{
		jobs <-j
		fmt.Println("sent job",j)
	}
	close(jobs)
	<-done
	
}