package main

import (
	"fmt"
	"time"
)

type Customer struct{
	name 	string
	balance int
}
var CustomerArr []Customer
func worker(wid int,jobs <-chan int,results chan<- Customer) {
	for j:= range(jobs) {
		fmt.Println("Creating Customer_",j," Balance=",j*100000)
		customer_name := fmt.Sprintf("Customer_%d",j)
		cust := Customer{customer_name,j*100000}
		time.Sleep(time.Second)
		results <- cust
	}
}

func makeTx(jobs <-Customer int,cr_de bool,amt int,results chan<- bool) {
	for j:= range(jobs){
		if cr_de {
			fmt.Println("Debiting Amount "+amt)
			j.balance = j.balance -amt
		} else {
			fmt.Println("Crediting Amount "+amt)
			j.balance = j.balance -amt
		}
		time.Sleep(1)
	}
}

func main() {
	const numCustomers = 20
	jobs := make(chan int,numCustomers)
	results := make(chan Customer,numCustomers)


	for w:=1;w<=10;w++ {
		go worker(w,jobs,results)
	}

	for j:=1;j<=numCustomers;j++ {
		jobs <-j
	}
	close(jobs)

	for a:=1;a<=numCustomers;a++ {
		CustomerArr = append(CustomerArr,<-results)
	}

	fmt.Println(CustomerArr)

	// numConcTrans = 5
	// txJobs := make(chan Customer,numConcTrans)
	// 
	


}