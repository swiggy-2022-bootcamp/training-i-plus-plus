package employee

import "fmt"

type Employee struct {
	Name         string
	Age          int
	Salary       int
	Designation  string
	Total_leaves int
	Leaves_taken int
}

func (e Employee) Remaing_leaves() {
	fmt.Println("Remaining leaves are:", e.Total_leaves-e.Leaves_taken)
}
