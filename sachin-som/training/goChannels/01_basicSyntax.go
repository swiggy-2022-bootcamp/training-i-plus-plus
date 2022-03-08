// var identifier chan datatype

package main

func main() {
	// declare a channel
	var ch1 chan string

	// Channels are reference type, so we need
	// to use make() for memory allocation

	ch = make(chan string)

	// short form
	// ch2 := make(chan int)

	// channel of channels
	// ch3 := make(chan chan int)

	// channel of funcs
	// funcChn := make(chan func)
}
