package employee

import "fmt"

type Employee struct {
	Name        string
	TotalLeaves int
	LeavesTaken int
}

func (e Employee) LeavesRemaining() {
	fmt.Println(e.Name, "has", e.TotalLeaves-e.LeavesTaken, "leaves remaining")
}
