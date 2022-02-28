/*
In this channel pattern, goRoutine processes what it receives
from an input channel and sends this to an output channel
*/

package main

import "fmt"

// generate will generate prime number checkers
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// filter would filter given channel using prime checker
func filter(in <-chan int, out chan<- int, prime int) {
	for {
		n := <-in
		if n%prime != 0 {
			out <- n
		}
	}
}

func main() {
	checkers := make(chan int)
	go generate(checkers)
	for {
		prime := <-checkers
		fmt.Println(prime)
		primes := make(chan int)
		go filter(checkers, primes, prime)
		checkers = primes
	}
}

// Output :- 2, 3, 5, 7, ...
