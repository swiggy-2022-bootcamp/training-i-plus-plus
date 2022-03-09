package employee

import "fmt"

type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func (e Employee) LeavesRemaining() {
	fmt.Println(e.TotalLeaves - e.LeavesTaken)
}
