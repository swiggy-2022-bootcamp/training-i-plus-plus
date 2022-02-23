package main 

import (
	"fmt"
	"time"
)

func display(msg string) {
	for i :=0;i <10;i++ {
		fmt.Println(msg,"   :",i)
		if(msg == "hello") {
			time.Sleep(time.Second*2)
		} else {
			time.Sleep(time.Second*1)
		}
	}
}

func main() {
	go display("Go routine")
	display("hello")

	go func(name string) {
		fmt.Println("My name is :",name)
	}("John")

	time.Sleep(time.Second*30)
	fmt.Println("end")
}