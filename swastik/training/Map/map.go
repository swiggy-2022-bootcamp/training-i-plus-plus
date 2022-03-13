package main

import (
    "fmt"
)

func main() {
	var mymap = make(map[int]string)

	//adding key value
	mymap[12] = "twelve"
	mymap[3] = "three"

	fmt.Println(mymap)

	map_2 := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	fmt.Println(map_2)

	//iteration
	for key, value := range map_2{
		fmt.Println("key: ", key, " value: ", value)
	}

	//retriving value
	fmt.Println(map_2[1])

	//removing element
	delete(map_2,1)
	fmt.Println(map_2)
}
