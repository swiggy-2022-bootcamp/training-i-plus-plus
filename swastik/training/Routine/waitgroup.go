package main 

import (
	"fmt"
	"sync"
)

func display(msg string){
	for i := 0 ; i < 10 ; i++ {
		fmt.Println(msg, " : " , i)
	}
}

func main(){
	var waitg sync.WaitGroup
	waitg.Add(1)
	go func(name string){
		defer waitg.Done()
		go display("Goroutine")
	}("John")


	waitg.Add(1)
	go func(name string){
		defer waitg.Done()
		fmt.Println("My name is : ",name)
	}("John")
	
	display("hello")
	waitg.Wait()
	fmt.Println("end")
}