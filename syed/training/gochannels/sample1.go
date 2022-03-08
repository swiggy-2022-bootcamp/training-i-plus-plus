package main

import (
	"fmt"
	
)

func main(){
	msgs := make(chan string)
 
	go func(){
		msgs <- "hello"
	}()
	
		
   res := <-msgs
	 fmt.Println(res)

}