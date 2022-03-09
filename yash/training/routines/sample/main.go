package main

import "sample/employee"

func main() {
	e := employee.Employee{
		Name:         "emp",
		Age:          25,
		Salary:       1000,
		Designation:  "manager",
		Total_leaves: 10,
		Leaves_taken: 10,
	}
	e.Remaing_leaves()
}
