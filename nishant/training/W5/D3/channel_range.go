package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)

	jobs <- 2
	jobs <- 45

	close(jobs)

	for j := range jobs {
		fmt.Println(j)
	}

	fmt.Println("Done")
}
