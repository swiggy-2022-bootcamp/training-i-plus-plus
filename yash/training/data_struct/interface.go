package main

import "fmt"

type calculate interface {
	rate_of_interest() float64
	processing_fee() float64
}

type principal_amount struct {
	amount float64
}

func (pa principal_amount) rate_of_interest() float64 {
	return pa.amount * 100
}

func main() {
	res := principal_amount{560}
	fmt.Println(res.rate_of_interest())
}
