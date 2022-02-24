package main

import "fmt"

func main(){
	currency := map[string]map[string]int{
		"Great Britain Pound":{"GBP":1},
		"Euro": {"EUR":2},
		"USA Dollar": {"USD":3},
	}

	for key, value := range currency{
		fmt.Printf("Currency Name: %v\n", key)

		for k, v := range value{
			fmt.Printf("\t Currency Code: %v\n\t\t\t Ranking: %v\n\n",k,v)
		}
	}
}