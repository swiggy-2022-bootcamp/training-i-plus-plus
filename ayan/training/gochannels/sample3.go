package main

import (
	"fmt"
	"time"
)

func task(t chan bool) {

	fmt.Print("Starting... ")
	time.Sleep(time.Second)
	fmt.Print("done... ")

	t <- true
}

func main() {
	t := make(chan bool, 1)

	go task(t)

	fmt.Print(<-t)
}
