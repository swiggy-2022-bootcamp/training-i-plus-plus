package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("temp.txt")
	if err != nil {
		fmt.Println("Error while opening the file: %v", err)
		return
	}
	// Close the file
	defer file.Close()

	// Create a buffer input reader
	var inputReader *bufio.Reader = bufio.NewReader(file)

	// Create buffer for reading from fd
	for {
		// buf, err := inputReader.ReadString('\n')
		buf, _, err := inputReader.ReadLine()
		if err == io.EOF {
			return
		}
		// fmt.Println("%s", buf)
		fmt.Printf("%s\n", buf)
	}
}
