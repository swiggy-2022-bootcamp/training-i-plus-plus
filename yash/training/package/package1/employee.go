package package1

import "fmt"

func PrintHello() {
	fmt.Println("Hello, Modules! This is mypackage speaking!")
}

type Employee struct {
	Name         string
	Age          int
	Salary       int
	Designation  string
	Leaves_taken int
	Total_leaves int
}

func (e Employee) Remaing_leaves() {
	fmt.Println("Remaining leaves are:", e.Total_leaves-e.Leaves_taken)
}
