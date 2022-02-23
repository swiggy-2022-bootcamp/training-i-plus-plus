package main

import (
	"fmt"
	"time"
)

func main(){
	c1 := make(chan string, 1)
	go func (){
		fmt.Println("inside go func 1")
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// fmt.Println("entering select 1")
	// select{
	// case res := <-c1:
	// 	fmt.Println(res)
	// case <-time.After(1 * time.Second):
	// 	fmt.Println("timeout 1")
	// }

	c2 := make(chan string, 1)
	go func(){
		fmt.Println("inside go func 2")
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	// fmt.Println("entering select 2")
	// select{
	// case res := <-c2:
	// 	fmt.Println(res)
	// case <-time.After(3 * time.Second):
	// 	fmt.Println("timeout 2")
	// }

	for i:=0; i<2; i++ {
		select{
		case msg1 := <-c1:
			fmt.Println("received msg1 -> ", msg1)
		case msg2 := <-c2:
			fmt.Println("received msg2 -> ", msg2)
		}
	}
}