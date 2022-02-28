package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("writeText.txt", d1, 0644)
	check(err)
	f, err := os.Create("writeText.txt")
	check(err)
	defer f.Close()
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)
	n3, err := f.WriteString(("writes\n"))
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)
	f.Sync()
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()
}
