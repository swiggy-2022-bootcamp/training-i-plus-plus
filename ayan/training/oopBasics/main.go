package main

import (
	emp "oopBasics/employee"
)

func main() {

	e := emp.Employee{
		Name:        "Ayan",
		TotalLeaves: 30,
		LeavesTaken: 10,
	}

	e.LeavesRemaining()

}
