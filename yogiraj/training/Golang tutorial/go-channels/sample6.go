package main

import (
	"fmt"
)

func main(){
	jobs := make(chan int,5)
	done := make(chan bool)

	go func(){
		for{
			j, more:=<-jobs
			if more{
				fmt.Println("received job",j)
			} else{
				fmt.Println("All jobs received")
				done <- true
				return
			}
		}
	}()

	for j :=1; j<=3;j++{
		jobs<-j
		fmt.Println("Sent job ", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}