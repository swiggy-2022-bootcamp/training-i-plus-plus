package main

import (
	"fmt"
	"strconv"
)

func main() {
	mp := make(map[string]int, 5)
	inv := make(map[int]string, 5)

	for i := 0; i < 5; i++ {
		mp["Item-"+strconv.Itoa(i)] = i
	}

	// Check if item present
	if _, present := mp["Item-1"]; present {

		// Delete an item
		delete(mp, "Item-1")
		fmt.Println(mp)
	} else {
		fmt.Println("Not present.")
	}

	// Inverting a map
	for key, Value := range mp {
		inv[Value] = key
	}

	fmt.Println(mp)
	fmt.Println(inv)

	// Creaing a slice of map
	sliceMap := make([]map[string]int, 5)

	for i := 0; i < 5; i++ {
		sliceMap[i] = make(map[string]int, 1)
		sliceMap[i] = map[string]int{"Slice-" + strconv.Itoa(i): i}
	}
	fmt.Println(sliceMap)
}

/*
map[Item-0:0 Item-2:2 Item-3:3 Item-4:4]
map[Item-0:0 Item-2:2 Item-3:3 Item-4:4]
map[0:Item-0 2:Item-2 3:Item-3 4:Item-4]
[map[Slice-0:0] map[Slice-1:1] map[Slice-2:2] map[Slice-3:3] map[Slice-4:4]]
*/
