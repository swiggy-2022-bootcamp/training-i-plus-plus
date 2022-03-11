package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	//printing
	fmt.Println("hello", "world")

	//vanilla variable declaration. Default val is 0 for int
	var num int = 10
	fmt.Println(num)

	//shorthand variable declaration - type inference
	a := 10.0
	b := 20.2
	sum := a + b
	fmt.Println(sum)

	//conditionals - else if or else should be preceded with closing bracket in same line, as compiler inserts ";" after predifined literals
	x := 10
	if x >= 10 {
		fmt.Println("Greater than 10")
	} else if x >= 5 {
		fmt.Println("Greater than 5")
	} else {
		fmt.Println("Less than 5")
	}

	//arrays
	var arr1 [5]int
	arr1 = [5]int{1, 2, 3, 4}
	fmt.Println(arr1)

	//short hand array declaration and init
	arr2 := [5]int{4, 3, 2, 1}
	fmt.Println(arr2)

	//dynamic array / slices - append will copy, then append and reassign new array
	arr3 := []int{1, 2, 3, 4, 5}
	arr4 := []int{}
	//arr4 is affected, arr3 isn't
	arr4 = append(arr3, 6)
	fmt.Println("arr3: ", arr3, " arr4: ", arr4)

	//loop through arr
	fmt.Println("For Loop Starts")
	for i := 0; i < len(arr4); i++ {
		fmt.Println(i, ", ", arr4[i])
	}
	fmt.Println("For Loop Ends")

	//Map data structure
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2

	delete(numbers, "one")

	for key, val := range numbers {
		fmt.Println(key, "->", val)
	}

	//func invocation
	fmt.Println("sum of 10, 29: ", getSum(10, 29))

	val, err := sqrt(-10)
	if err != nil {
		fmt.Println("Sqrt of -10: ", err)
	} else {
		fmt.Println("Sqrt of 10:", val)
	}

	//custom type - struct
	var student1 student = student{name: "Ravi", age: 22, marks: 88}
	fmt.Println(student1)

	//pointers
	var number = 1
	increment(&number)
	fmt.Println(number)

}

func increment(num *int) {
	//dereference
	*num++
}

type student struct {
	name  string
	age   int
	marks float64
}

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("Negative input is unacceptable")
	}
	return math.Sqrt(num), nil
}

func getSum(a int, b int) int {
	return a + b
}
