package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var FileFlag = flag.Bool("file", false, "Name of the file, which would be read.")

func main() {

	// parse args
	flag.Parse()

	// take out each filename given in cmd
	for i := 0; i < flag.NArg() && *FileFlag; i++ {
		file, err := os.Open(flag.Arg(i))
		if err != nil {
			panic(err)
		}
		defer file.Close()
		catFile(bufio.NewReader(file))

	}
}

func catFile(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			return
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
	return
}

// CMD : go run 09_practice.go -file temp.txt
