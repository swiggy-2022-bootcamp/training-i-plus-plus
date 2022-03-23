package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "try channel 1"
    }()
    go func() {
        time.Sleep(5 * time.Second)
        c2 <- "try channel 2"
    }()

    for i := 0; i < 2; i++ {
        select {
            case msg1 := <-c1: // Received in Channel 1
                fmt.Println("received", msg1)
            case msg2 := <-c2: // Received in channel 2
                fmt.Println("what", msg2)
        }
    }
}