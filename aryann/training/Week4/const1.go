package main

import "fmt"

func main() {

	const name, country = "aryann", "india"

	const (
		company_name string  = "swiggy"
		salary       float64 = 100.21
	)

	fmt.Println(company_name, salary)
	fmt.Println(name, country)
}
