package main

import "fmt"

func msg(name string)(m string){
	m = "Welcome, " + name
	return 
}

func main() {
	var m string
	m = msg("John")
	fmt.Println(m)
}