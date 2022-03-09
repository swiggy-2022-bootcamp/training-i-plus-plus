package main

import (
	"fmt"
	"time"
)

func numbers() {

	for i := 1; i <= 50; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Print(" ", i)
	}
}

func alphabets() {

	for i := 'a'; i <= 'z'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf(" %c", i)
	}
}

func main() {
	go numbers()
	go alphabets()
	time.Sleep(9000 * time.Millisecond)
	fmt.Println(" End of main()!!!")
}
