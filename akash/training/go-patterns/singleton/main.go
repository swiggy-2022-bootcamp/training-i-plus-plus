package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

//var lock = &sync.Mutex{}

type Database struct{}

var instance *Database

func getInstance() *Database {
	wg.Add(1)
	defer wg.Done()
	//if instance == nil {
	instance = &Database{}
	//}
	fmt.Printf("returning instance : %p\n", instance)
	return instance
}

func main() {
	for i := 0; i < 10; i++ {
		go getInstance()
		//time.Sleep(time.Second * 5)
	}
	wg.Wait()
}
