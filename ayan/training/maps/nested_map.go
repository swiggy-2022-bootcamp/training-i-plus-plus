package main

import "fmt"

func main() {

	currency := map[string]map[string]int{
		"UK": {"GBP": 1},
		"EU": {"EUR": 2},
		"US": {"USD": 3},
	}

	for key, value := range currency {
		fmt.Println("Country:", key)
		for k, v := range value {
			fmt.Println("\t Ccy Code:", k)
			fmt.Println("\t\t Ranking:", v)
		}
	}
}
