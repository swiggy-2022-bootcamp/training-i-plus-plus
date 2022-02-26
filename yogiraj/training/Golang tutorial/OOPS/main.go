package main

import "OOPS/employee"

func main(){
	e := employee.Employee{
		FirstName: "Yogiraj",
		LastName: "Gutte",
		TotalLeaves: 30,
		LeavesTaken: 10,
	}

	e.LeavesRemaining()
}