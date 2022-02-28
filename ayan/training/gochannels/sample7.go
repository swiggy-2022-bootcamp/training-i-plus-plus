package main

import (
	"fmt"
)

func main() {

	queue := make(chan string, 4)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	queue <- "four"
	// close(queue)

	for i := 0; i < 2; i++ {
		fmt.Println(<-queue)
	}

	// for elem := range queue {
	// 	fmt.Println(elem)
	// }

}
