package main

import (
	trs "TrainReservationSystem/Ticket"
	"fmt"
	"sync"
	"time"
)

func ticketBookingWithConcurrency(wg *sync.WaitGroup, firstName string, lastName string, age int) {
	defer wg.Done()
	ticket := trs.Ticket{}
	passenger1 := trs.Passengers{FirstName: firstName, LastName: lastName, Age: age}

	//passengerList := make([]Passengers, 2)
	//passengerList[0] = passenger1
	//passengerList[1] = passenger2

	//var passengerList []Passengers
	passengerList := []trs.Passengers{}
	passengerList = append(passengerList, passenger1)

	ticket.SetPassengers(passengerList)
	ticket.SetSource("Delhi")
	ticket.SetDestination("Mumbai")
	ticket.SetAmount()

	time.Sleep(2 * time.Second)

	fmt.Printf("%+v\n", ticket)

}

func ticketBookingWithoutConcurrency(firstName string, lastName string, age int) {
	ticket := trs.Ticket{}
	passenger1 := trs.Passengers{FirstName: firstName, LastName: lastName, Age: age}

	//passengerList := make([]Passengers, 2)
	//passengerList[0] = passenger1
	//passengerList[1] = passenger2

	//var passengerList []Passengers
	passengerList := []trs.Passengers{}
	passengerList = append(passengerList, passenger1)

	ticket.SetPassengers(passengerList)
	ticket.SetSource("Delhi")
	ticket.SetDestination("Mumbai")
	ticket.SetQuota(&ticket, "Tatkal Quota")

	ticket.SetAmount()

	time.Sleep(2 * time.Second)

	fmt.Printf("%+v\n", ticket)

}

func ticketBookingUsingChannels(ticketBookingChannel chan trs.Ticket, firstName string, lastName string, age int) {
	ticket := trs.Ticket{}
	passenger1 := trs.Passengers{FirstName: firstName, LastName: lastName, Age: age}

	//passengerList := make([]Passengers, 2)
	//passengerList[0] = passenger1
	//passengerList[1] = passenger2

	//var passengerList []Passengers
	passengerList := []trs.Passengers{}
	passengerList = append(passengerList, passenger1)

	ticket.SetPassengers(passengerList)
	ticket.SetSource("Delhi")
	ticket.SetDestination("Mumbai")
	ticket.SetQuota(&ticket, "Tatkal Quota")
	ticket.SetAmount()

	time.Sleep(1 * time.Second)

	ticketBookingChannel <- ticket

	fmt.Printf("%+v\n", ticket)
}

func main() {

	//// scenario - 1
	//startTime := time.Now()
	//ticketBookingWithoutConcurrency("Taranjeet", "Singh", 24)
	//ticketBookingWithoutConcurrency("Manjeet", "Singh", 48)
	//endTime := time.Now()
	//timeTaken := endTime.Sub(startTime)
	//
	//fmt.Printf("Time taken - %v\n", timeTaken)
	//
	//// scenario - 2
	//
	//startTime = time.Now()
	//var wg sync.WaitGroup
	//wg.Add(2)
	//go ticketBookingWithConcurrency(&wg, "Taranjeet", "Singh", 24)
	//go ticketBookingWithConcurrency(&wg, "Manjeet", "Singh", 48)
	//endTime = time.Now()
	//timeTaken = endTime.Sub(startTime)
	//
	//wg.Wait()
	//
	//fmt.Printf("Time taken - %v\n", timeTaken)

	//scenario - 3

	ticketBookingChannel := make(chan trs.Ticket)
	go ticketBookingUsingChannels(ticketBookingChannel, "Taranjeet", "Singh", 24)
	go ticketBookingUsingChannels(ticketBookingChannel, "Manjeet", "Singh", 48)

	select {
	case ticket := <-ticketBookingChannel:
		fmt.Printf("%+v\n", ticket)
	}
}
