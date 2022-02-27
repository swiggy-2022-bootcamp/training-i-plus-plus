package main

import "fmt"

func forLoopDemo() {
	for d := 1; d <= 7; d++ {
		fmt.Printf("%v, ", d)
	}
	fmt.Println()
}

func iterateOverAnArrayUsingForLoopRange(days []string) {
	for _, d := range days {
		fmt.Printf("%v | ", d)
	}
	fmt.Println()
}

// the closest GoLang can get to for loops ;)
func conditionalForLoopDemo() {
	i := 0
	for i < 5 {
		fmt.Print(i, " ")
		i += 1
	}
	fmt.Println()
}

func main() {
	fmt.Println("Welcom to loops in Go")

	daysArr := []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}
	forLoopDemo()
	iterateOverAnArrayUsingForLoopRange(daysArr)
	conditionalForLoopDemo()
}
