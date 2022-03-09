package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func display1(str string) {
	defer wg.Done()

	for i := 0; i < 50; i++ {
		fmt.Println(str)
	}
}

func display2(str string) {
	defer wg.Done()

	for i := 0; i < 50; i++ {
		fmt.Println(str)
	}
}

func main() {

	go display1("hello")

	go display2("world")

	time.Sleep(time.Millisecond * 200)

	wg.Add(2)

	go display1("go")

	go display2("routine")

	wg.Wait()

}
