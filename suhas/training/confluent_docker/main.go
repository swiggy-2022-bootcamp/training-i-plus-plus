package main

import "time"

func main() {
	go produce()
	go consume()
	time.Sleep(2 * time.Minute)
}
