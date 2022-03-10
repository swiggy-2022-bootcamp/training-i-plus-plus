package employee

import "fmt"

type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	TakenLeaves int
}

func CalRemaining(e *Employee) {
	fmt.Println("Remaining Leaves : ", e.TotalLeaves-e.TakenLeaves)
}
