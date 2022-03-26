package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func randSleep(wg *sync.WaitGroup, function string, length int) {
	for i := 1; i <= length; i++ {
		fmt.Println(function, rand.Intn(i))
		time.Sleep(time.Second)
	}
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	fmt.Println("Without Go Routine")
	withoutGoRoutine(wg)
	fmt.Println("With Go Routine")
	withGoRoutine(wg)
	wg.Wait()
}

func withoutGoRoutine(wg *sync.WaitGroup) {
	randSleep(wg, "first:", 5)
	randSleep(wg, "second:", 5)
}

func withGoRoutine(wg *sync.WaitGroup) {
	go randSleep(wg, "first:", 5)
	go randSleep(wg, "second:", 5)
}
