package trs

import "fmt"

type Ticket struct {
	passengersName []Passengers
	source         string
	destination    string
	amount         int
	seatNumbers    []int
	distance       int
	quota          Quota
	train          Train
}

type Seat struct {
	seatNumber int
	seatType   string
}

const farePerKilometer int = 100

func (t *Ticket) SetSource(source string) {
	t.source = source
}

func (t *Ticket) SetDestination(destination string) {
	y := Passengers{FirstName: "tara", LastName: "dqw", Age: 23}
	x := make([]Passengers, 0)
	x = append(x, y)
	t.SetPassengers(x)
	t.destination = destination
}

func (t *Ticket) getAmount() int {
	fmt.Printf("+%v\n", t.quota)
	amount := t.quota.CalculateFare() + t.distance*farePerKilometer
	return amount
}

func (t *Ticket) SetAmount() {
	amount := t.getAmount()
	t.amount = len(t.passengersName) * amount
}

func (t *Ticket) SetQuota(ticket *Ticket, name string) {
	quota, err := CreateQuotaFactory(name)
	if err == nil {
		ticket.quota = quota
	}

}

func (t *Ticket) GetQuota() Quota {
	return t.quota

}
