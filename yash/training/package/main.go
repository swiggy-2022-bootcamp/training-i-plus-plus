package main

import "package/package1"

func main() {
	e := package1.Employee{
		Name:         "emp",
		Age:          25,
		Salary:       1000,
		Designation:  "manager",
		Leaves_taken: 10,
		Total_leaves: 20,
	}
	e.Remaing_leaves()
}
