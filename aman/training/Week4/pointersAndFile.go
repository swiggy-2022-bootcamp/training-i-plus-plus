package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	item := ""
	passByValue(item)
	fmt.Println(item)
	passByReference(&item)
	fmt.Println(item)

	readFile()
}

func passByValue(item string) {
	item = "text by value"
}

func passByReference(item *string) {
	*item = "text by reference"
}

func readFile() {
	content, err := ioutil.ReadFile("text.txt") //Path to file, default in code directory
	if err == nil {
		fmt.Println(string(content))
	} else {
		fmt.Println("Error: ", err)
	}
}
