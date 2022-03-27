package main

import (
	"fmt"
	"sync"
	"time"
)

func workerfunc(id int){
	fmt.Println("started worker ",id)
	time.Sleep(time.Second)
	fmt.Println("finished worker ",id)
}

func main(){
	var wg sync.WaitGroup
	for i:=1;i<=10;i++{
		wg.Add(1)
		i :=i
		go func(){
			defer wg.Done()
			workerfunc(i)
		}()
	}

	wg.Wait()
}