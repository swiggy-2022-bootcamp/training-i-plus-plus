package main

import (
	"fmt"
	"time"
)

func display(s string){
	for i:=1; i<50; i++ {
		fmt.Printf(" i => %d, %s\n", i,s)
	}
}
func main(){
	go display("Go-Routine")
	go func(name string){
		fmt.Println("My name is : ",name)
	}("John")
	display("Hello")
	time.Sleep(time.Second)
	fmt.Println("end")
}
