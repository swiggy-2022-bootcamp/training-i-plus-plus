package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	d1 := []byte("writing to file from go")
	os.WriteFile("./write1.txt", d1, 0644)

	f, _ := os.Create("./write2.txt")
	defer f.Close()

	d2 := []byte{123, 99, 105, 113}
	f.Write(d2)

	f.WriteString("some string")

	w := bufio.NewWriter(f)
	nw, _ := w.WriteString("from buff")
	fmt.Println("nw", nw)
	w.Flush()
}
