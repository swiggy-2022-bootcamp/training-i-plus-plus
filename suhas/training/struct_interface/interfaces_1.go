package main 

import "fmt"

type calculate interface {
	rate_of_interest() float64
	processing_fees() float64
}

type prinicpal_amt struct {
	amount float64
}

func (pa prinicpal_amt)rate_of_interest() float64 {
	return pa.amount * 0.01
}

func main() {
	var result calculate
	result = prinicpal_amt{560}
	fmt.Println(result.rate_of_interest())
}