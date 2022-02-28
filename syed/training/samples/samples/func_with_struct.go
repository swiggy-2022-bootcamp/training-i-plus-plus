package main

import "fmt"

type animal struct {
	name            string
	characteristics []string
}

func main() {

	animal1 := animal{
		name: "Elephant",
	}

	animal1.run()
}

func (a animal) run() {
	fmt.Println(a.name, "is a lazy animal hence cannot run")
}