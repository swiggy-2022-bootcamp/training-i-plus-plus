package main

import "oop/employee"

func main() {

	e := employee.Employee{
		FirstName:   "aryann",
		LastName:    "dhir",
		TotalLeaves: 30,
		LeavesTaken: 10,
	}

	e.LeavesRemaining()
}
