package main

import "fmt"

type animal struct {
	name             string
	charachteristics []string
}
type vertex struct {
	X int
	Y int
}

func main() {
	// animal := animal{
	// 	name: "dog",
	// 	charachteristics: []string{"likes to eat", "likes to sleep",
	// 		"likes to play"},
	// }
	// fmt.Println(animal.name)
	// for _, v := range animal.charachteristics {
	// 	fmt.Printf("\t %v\n", v)

	// }
	temp := animal{}
	temp.name = "hen"
	fmt.Println(temp.name)
	v := vertex{1, 2}
	fmt.Println(v)
	animal1 := animal{}
	animal1.name = "hen"
	animal1.run()
}
func (a animal) run() {
	fmt.Println(a.name, " animal is running")
}
