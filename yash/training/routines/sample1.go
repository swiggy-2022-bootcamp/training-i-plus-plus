package main

import (
	"fmt"
	"time"
)

func display(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg, " : ", i)
		// display("hello")
	}
}
func main() {
	// display("hello")
	for i := 0; i < 10; i++ {
		s := fmt.Sprintf("%d", i)
		temp := "Go Rountines " + s
		go display(temp)
	}
	go display("GO Rotuines")
	go func(name string) {
		fmt.Println("The name is : ", name)
	}("John" + time.Now().String())
	display("hello")
	time.Sleep(time.Second * 2)
	fmt.Println("The main function is over")

}
