package main

import (
	"fmt"
	"time"
)


func display(x string){
	fmt.Println("hello there ",x)
}

func main(){	
	display("pradyuman")

	go display("swapnil")

	go func (x string){
		fmt.Println("inside anonymous function ",x)
	}("shilpi")
	
	time.Sleep((time.Second))
}