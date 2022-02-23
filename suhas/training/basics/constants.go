package main

import "fmt"

func main() {
	const name, country = "John","India"

	const (
		company_name string  = "IBM"
		salary 		 float64 =  5700064
	)
	
	fmt.Println(company_name,salary)
	fmt.Println(name,country)
}