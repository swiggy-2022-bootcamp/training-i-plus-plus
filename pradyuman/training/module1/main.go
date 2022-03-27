package main

import "module1/employee"

func main(){
	w:= employee.Worker{
		Name: "pradyuman",
		Age : 21,
	}

	w.HelloWorker()
}