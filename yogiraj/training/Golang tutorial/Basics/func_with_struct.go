package main

import "fmt"

type animal struct {
	name            string
	characteristics []string
}

func main() {
	animal1 := animal{
		name:            "Lion",
		characteristics: []string{"Eats animal", "Wild animal", "King of jungle"},
	}
	animal1.run()
	println(animal1.name)
}

func (a animal) run(){
	fmt.Println(a.name, "has characteristics: ", a.characteristics)
	
}