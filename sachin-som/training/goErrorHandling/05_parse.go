package main

import (
	"fmt"

	"github.com/sachinsom93/training-i-plus-plus/sachin-som/training/goErrorHandling/myParse"
)

var (
	example = []string{
		"1 2 3 4 5",
		"2 + 2 = 4",
		"1st class",
		"",
	}
)

func main() {
	for _, str := range example {
		fmt.Printf("Parsing %q:\n", str)
		nums, err := myParse.Parse(str)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(nums)
	}
}
