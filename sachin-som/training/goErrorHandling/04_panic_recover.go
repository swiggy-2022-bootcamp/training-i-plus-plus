package main

import "fmt"

func badCall() {
	panic("bad end")
}

func test() {
	// regain control of a panicking goroutine
	// recover only usefull when called inside a deffered function
	go func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicking %s\n", e)
		}
	}()
	badCall() // recover will stop the remaining sequence
	fmt.Println("After Bad call.")
}
func main() {
	fmt.Println("Calling test")
	test()
	fmt.Println("test completed")
}

// Summary - panic causes the stack to unwind untill a deffered function
// is found or the program terminates
