package trs

type Passengers struct {
	FirstName string
	LastName  string
	Age       int
}

func (t *Ticket) SetPassengers(passengersName []Passengers) {
	t.passengersName = passengersName

}
