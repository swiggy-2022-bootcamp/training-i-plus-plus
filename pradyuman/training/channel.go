package main

import (
	"fmt"
	"sync"
	"time"
)

func workerfunc(id int, messages chan string){
	fmt.Println("started worker ",id)
	messages<-"hello"
	fmt.Println("mssg sent")
	time.Sleep(time.Second)

	fmt.Println("finished worker ",id)
}

func main(){
	var wg sync.WaitGroup

	messages:= make(chan string,3)

	for i:=1;i<=1;i++{
		wg.Add(1)
		i :=i
		go func(){
			defer wg.Done()
			workerfunc(i,messages)
		}()
	}
	fmt.Println("here")
	time.Sleep(time.Second)
	fmt.Println("here2")
	msg := <-messages
	fmt.Println(msg)
	wg.Wait()
}