package main

import "fmt"

func main() {

	s := make([]int, 4)
	s[0] = 10
	s[1] = 20
	s[2] = 30
	s[3] = 40
	fmt.Println("Slice created with 'make': ", s)

	//a. We can add elements to this slice using "append" builtin function:
	s = append(s, 50)
	fmt.Println("Added one element to slice: ", s)

	//b. We can append more than one element to the slice:
	s = append(s, 60, 70)
	fmt.Println("Added two elements to slice: ", s)

	//c. We can remove from that slice:
	//Say we want to remove "30" which is the third element at index "2"
	s = append(s[:2], s[2+1:]...) //... is a variadic argument in Go
	fmt.Println("Deleted one element from slice: ", s)

	//d. Replace an element with another
	// Say i want replace the third element now "40" with the fifth element "60":
	s[2] = s[len(s)-2]
	fmt.Println("Slice with element replaced: ", s)

	// Say i want replace the third element now "60" with the last element "70":
	s[2] = s[len(s)-1]
	fmt.Println("Updated Slice with element replaced: ", s)

	//e. Get particular elements from the slice:
	//We now have this slice: [10 20 70 50 60 70]
	//To get the 2nd(index 1) to the 4th(index 3) element, we do:
	s = s[1:4]
	fmt.Println("Slice with second to fourth element: ", s)

	//f. Get the length of the current slice:
	fmt.Println("Length: ", len(s))

	//g. Get the capacity of the current slice:
	fmt.Println("Capacity: ", cap(s)) //this is give "7"

	//h. Copy one slice to another:
	//Make a slice with the same length as "s":
	d := make([]int, len(s))
	copy(d, s)
	fmt.Println("This is the new slice: ", d)
}