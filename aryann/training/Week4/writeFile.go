package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	f, err := os.Create("/tmp/dat2")
	check(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)

	fmt.Println("wrote", n2, "bytes")
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Println("wrote", n3, "bytes")

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Println("wrote", n4, "bytes")

	w.Flush()
}
