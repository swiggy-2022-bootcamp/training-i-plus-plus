package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// os.Args --> Slice of strings
	// os.Args[0] --> build path of exe file
	// fmt.Println(os.Args)
	who := ""
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Printf("Welcome %s\n", who)
}
