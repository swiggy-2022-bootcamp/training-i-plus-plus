package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("hello")

	file, err := os.Create("./sample.txt")

	if err != nil {
		panic(err)
	}

	content := "Is this just fantasy"
	io.WriteString(file, content)
	defer file.Close()

	readFile("sample.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(databyte))
}
