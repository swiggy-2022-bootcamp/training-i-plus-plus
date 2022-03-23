package main

import "fmt"

func main() {
	//catching panics in caller function via async communication using channels
	err := make(chan error)
	ok := make(chan bool)
	go panicReplacement(err, ok)
	if !<-ok {
		fmt.Println("error resolved: ", <-err)
	}

	//go spreadPanic()

	//wait indefinitely

}

//Most recent panic will replace it's predecessor panics.
//In this example, at any time there is only one panic
func panicReplacement(err chan error, ok chan bool) {
	defer func() {
		//panic(3) shall be recovered
		if r := recover(); r != nil {
			_, ok1 := r.(error) //type assertion
			ok <- ok1
			if !ok1 {
				err <- fmt.Errorf("error: %v", r)
			}
		}
	}()

	//will replace panic(2)
	defer panic(3)

	//will replace panic(1)
	defer panic(2)

	panic(1)
}

//At a time there can be atmost 2 unrecovered panics.
//when a panic occurs in nested func, it spreads to outer func when inner func (where panic occured) exits
func spreadPanic() {
	//3. After this deferred anonymous func exits, panic(3) replaces panic(1). In this particular example, panic(3) is unrecovered
	defer func() {
		//2. panic(3) replaces panic(2) as soon as deferred anonymous func exits
		defer panic(3)
		func() {
			//1. at this point, there are 2 unrecovered panics, coexisting. One is panic(1) of spreadPanic func. And the other is panic(2) of the anonymous func in deferred func
			//it wont stay the same for too long, bcs panics spread
			panic(2)
		}()
	}()
	panic(1)
}
