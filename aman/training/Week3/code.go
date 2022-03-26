package main

import "fmt"

func main() {
	fmt.Println("Go is working!")
	switchFunc()
	forLoop()
	whileLoopImplementation()
	rangeKeyWord()
	mapInGo()
	rectangle := Rectangle{height: 3, width: 5}
	fmt.Println(rectangle.area())
}

func switchFunc() {
	income := 10000

	if income >= 5000 {
		fmt.Println("Credit card approved")
	} else {
		fmt.Println("Income insufficent")
	}

	switch income {
	case 5000:
		fmt.Println("Income is less")
	case 10000:
		fmt.Println("Income is enough")
	default:
		fmt.Println("Income not in options")
	}
}

func forLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func whileLoopImplementation() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i += 1
	}
}

func rangeKeyWord() {

	var array []int = []int{1, 2, 3, 4, 5}
	for i, j := range array {
		fmt.Println(i, j)
	}
}

func mapInGo() {
	courseMap := make(map[string]string)
	courseMap["Teacher1"] = "Computer Science"
	fmt.Println(courseMap)
	courseMap["Teacher2"] = "Electronics"
	fmt.Println(courseMap)
	delete(courseMap, "Teacher1")
	fmt.Println(courseMap)

}

type Rectangle struct {
	height float64
	width  float64
}

func (rect *Rectangle) area() float64 {
	return rect.width * rect.height
}
