package main

import (
	"flag"
	"fmt"
)

func main() {

	nFlag := flag.Bool("n", false, "This is help message") // Returns the address of bool var
	flag.Parse()                                           // Parsed all the flag commands from os.Args[1:], enables flag.Args() and flag.Arg(i)
	fmt.Println(flag.Args())                               // Non flag command line args

	var s string
	if *nFlag { // Checks if -n is used or not
		for i := 0; i < flag.NArg(); i++ {
			if i > 0 {
				s += " "
			}
			s += flag.Arg(i)
		}
	}
	fmt.Println(s)
	// fmt.Println(flag.Arg()) // ith flag command
}
