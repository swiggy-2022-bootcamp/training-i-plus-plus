package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	inputReader *bufio.Reader
	msg         string
	err         error
)

func main() {
	inputReader = bufio.NewReader(os.Stdin) // Returns a reader of default buf size (4096)
	fmt.Println("Enter some input data: ")
	buf, err := inputReader.ReadString('\n')
	if err != nil {
		panic("Error while taking input.")
	}
	fmt.Printf("Your data: %v", buf)
}
