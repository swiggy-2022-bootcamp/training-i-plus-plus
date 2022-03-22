package main

import (
	"fmt"
	"time"
)

type train struct {
	coach string
	src   string
	dst   string
	date  time.Time
}

func (user train) train_details() {
	fmt.Println("Seat type: ", user.coach)
	fmt.Println("From: ", user.src)
	fmt.Println("To: ", user.dst)
	fmt.Println("Date: ", user.date)
	fmt.Println("==========================================\n")
}

func main() {

	t1 := train{"AA", "Bangalore", "Chennai", time.Now()}
	t1.train_details()

	t2 := train{"AAA", "Mumbai", "Chennai", time.Now()}
	t2.train_details()
}
