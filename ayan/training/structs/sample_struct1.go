package main

import "fmt"

type animal struct {
	name string
	char []string
}

func main() {
	a1 := animal{
		name: "Lion",
		char: []string{
			"Carnivorous",
			"King of the jungle",
			"Wild animal",
		},
	}
	fmt.Println("Animal Name:", a1.name)
	for _, v := range a1.char {
		fmt.Println("\t", v)
	}

	fmt.Println(a1.run())
}

func (a animal) run() (s string) {
	return a.name + " is too lazy to run"
}
