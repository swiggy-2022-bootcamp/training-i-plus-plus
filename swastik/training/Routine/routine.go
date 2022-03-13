package main 

import (
	"fmt"
	"time"
)

func display(msg string){
	for i := 0 ; i < 500 ; i++ {
		fmt.Println(msg, " : " , i)
	}
}

func main(){
	
	go display("Goroutine")

	go func(name string){
		fmt.Println("My name is : ",name)
	}("John")

	display("hello")
	time.Sleep(time.Second)
	fmt.Println("end")
}