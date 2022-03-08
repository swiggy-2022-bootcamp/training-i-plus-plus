package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	// Map as a pointer
	map1 := new(map[string]int)

	// Map as a value type
	var map3 map[string]int

	// Map using make
	// can also mention cap
	map2 := make(map[string]int, 5)

	// Map as a literal
	map1 = &map[string]int{"One": 1, "Two": 2, "Three": 3}

	// Maps are referenced type
	map3 = map2

	map2["first"] = 1
	map2["Second"] = 2
	map2["Third"] = 3

	// Would alter map2 and map3
	map3["First"] = 11

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(len(map2))
	// fmt.Println(cap(map2))
	fmt.Println(map3)
}
