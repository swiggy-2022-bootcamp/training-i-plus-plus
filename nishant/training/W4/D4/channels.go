package main

import (
	"fmt"
)

func rout(ch chan int) {
	//ch <- 1003
	fmt.Println("inside rout")
	for {

	}
}

func main() {

	ch := make(chan int)

	go rout(ch)
	fmt.Println("waiting for receive")
	fmt.Println(<-ch)
	fmt.Println("done")

	ch2 := make(chan int, 5)

	ch2 <- 10
	ch2 <- 20
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	ch2 <- 30
	fmt.Println(<-ch2)

}
