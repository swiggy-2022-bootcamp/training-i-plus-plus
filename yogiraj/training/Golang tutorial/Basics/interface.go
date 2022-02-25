package main

import "fmt"

type calculate interface{
	rate_of_interest() float64
	// processing_fee() float64
}

type principal_amt struct{
	amount float64
}

func (pa principal_amt) rate_of_interest() float64{
	return pa.amount * 100
}

func main() {
	var res calculate
	res = principal_amt{560}
	fmt.Println(res.rate_of_interest())
}