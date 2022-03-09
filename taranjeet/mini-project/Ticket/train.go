package trs

type Train struct {
	trainNumber    int
	trainName      string
	origin         string
	destination    string
	seatsAvailable []int
	seatsReserved  []int
}

func (t *Train) getAvailableSeat(trainNumber int) []int {
	return t.seatsAvailable
}

func (t *Train) getReservedSeat(trainNumber int) []int {
	return t.seatsReserved
}

func (t *Train) getTrainName(trainNumber int) string {
	return t.trainName
}

func (t *Train) setAvailableSeats(seats []int) {
	t.seatsAvailable = seats
}

func (t *Train) setReservedSeats(seats []int) {
	t.seatsAvailable = seats
}

func (t *Train) bookSeat(trainNumber int, seats []int) {
	availableSeatMap := make(map[int]bool)
	availableSeat := make([]int, 0)

	for i := 0; i < len(t.seatsAvailable); i++ {
		availableSeatMap[t.seatsAvailable[i]] = true
	}

	for i := 0; i < len(seats); i++ {
		_, ok := availableSeatMap[seats[i]]
		if ok == false {
			availableSeat = append(availableSeat, seats[i])
		}
	}

	t.setAvailableSeats(availableSeat)
	t.setReservedSeats(append(t.seatsReserved, seats...))

}
