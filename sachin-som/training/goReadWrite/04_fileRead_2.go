package main

import (
	"fmt"
	"io/ioutil"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	inputFile := "temp.txt"
	outFile := "out.txt"

	// Read file
	content, err := ioutil.ReadFile(inputFile)
	checkErr(err)

	fmt.Printf("File content: \n  %v\n", string(content))

	// Write file
	writeErr := ioutil.WriteFile(outFile, content, 0x644)
	checkErr(writeErr)
}
