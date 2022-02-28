package main

import "fmt"

func msg(name string) (m string) {
	m = "Welcome, " + name
	return
}

func main() {
	m := msg("john")
	fmt.Println(m)
}
