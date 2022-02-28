package main

import "package_example/employee"

func main() {
	e := employee.Employee{
		FirstName:   "John",
		LastName:    "Cilofi",
		TotalLeaves: 30,
		LeavesTaken: 10,
	}
	e.LeavesRemaining()

}
