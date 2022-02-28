package main

import (
	"fmt"
	"time"
)

func channel_task(t chan bool) {
	fmt.Print("starting...")
	time.Sleep(time.Second)
	fmt.Print("done...")

	t <- true
}

func main() {
	t := make(chan bool, 1)
	go channel_task(t)

	fmt.Print(<-t)
	go channel_task(t)
	fmt.Print(<-t)

}
