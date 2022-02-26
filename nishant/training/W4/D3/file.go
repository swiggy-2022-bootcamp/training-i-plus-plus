package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	str, err := os.ReadFile("./test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(str))

	f, _ := os.Open("./test.txt")
	b1 := make([]byte, 4)
	n1, _ := f.Read(b1)
	fmt.Println(n1, " : ", b1, string(b1))

	b2 := make([]byte, 10)
	n2, _ := io.ReadAtLeast(f, b2, 5)
	fmt.Println(n2, " : ", b2, string(b2))

	br := bufio.NewReader(f)

	b4, _ := br.Peek(6)
	fmt.Println(b4, string(b4))
}
