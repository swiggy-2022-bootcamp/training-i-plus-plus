package main

import (
	"fmt"
    "os"
)

type passenger struct {
	fname string
	lname string
	passid int
}

type ticket struct {
	src string
	dest string
	cust passenger
}

type  disp interface {
	display()
	check() bool
}	

var id = 99
var tickets = map[string]int {
	"Banglore Mysore" :10,
	"Mysore Manglore" :10,
	"Chennai Banglore":10,
}

var passenger_array = []passenger{}

// main 
func main() {
	BookTicket("John","Doe","Banglore","Mysore")
	BookTicket("AB","CD","Delhi","Noida")
	BookTicket("WQ","LM","Chennai","Banglore")
	BookTicket("John","Doe","Chennai","Banglore")
}


// util functions
func getPassenger(fname string, lname string) passenger {
	id++
	return passenger{fname,lname,id}
}

func getTicket(ps passenger, source string , destination string) ticket {
	return ticket{source,destination,ps}
}

func (tk ticket)printTicket() {
	fmt.Println("Ticket Confirmed :)")
	fmt.Println("Source        : ",tk.src)
	fmt.Println("Destination   : ",tk.dest)
	fmt.Println("Passenger name: ",tk.cust.fname,tk.cust.lname)
	fmt.Println("Passenger ID  : ",tk.cust.passid)
	fmt.Printf("==============================\n\n")
}

func check(e error) {
	if e!= nil {
		panic(e)
	}
}

func (tk ticket)printTicketFile() {
	placrholderstring := `Source            : %s
Destincation      : %s
Name              : %s %s
ID                : %d`
	stringToWrite := fmt.Sprintf(placrholderstring,tk.src,tk.dest,tk.cust.fname,tk.cust.lname,tk.cust.passid+1)

	f, err := os.OpenFile("ticket.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(stringToWrite+"\n\n"); err != nil {
		panic(err)
	}
}

func (tk ticket)check() bool {
	var srcdest = tk.src + " "+ tk.dest
	v,exist := tickets[srcdest]
	if exist && v>0  {
		tickets[srcdest]-=1
		id-=1
		return true
	}
	return false
}

func (ps passenger) check() bool {
	for _, v := range passenger_array {
		if v.fname == ps.fname && v.lname == ps.lname{
			return true
		}
	}
	return false
}


func (ps passenger) display() {
	fmt.Printf("Passenger Details => firstname: %s lastname: %s id :%d\n",ps.fname,ps.lname,ps.passid)
}

func (tk ticket) display() {
	fmt.Printf("Ticket queried => Source: %s Destination: %s\n",tk.src,tk.dest)
}

func displaydetails(d []disp) {
	for _,k := range d{
		k.display()
	}
}


func BookTicket(fname string, lname string, src string ,dest string) {
	passenger := getPassenger(fname,lname)
	ticket := getTicket(passenger,src,dest)
	displaydetails([]disp{ticket,passenger})
	if(passenger.check()) {
		fmt.Println("Passenger already exist ~", passenger.fname,passenger.lname)
	} else {
		fmt.Println("New Passenger *",passenger.fname,passenger.lname)
		passenger_array = append(passenger_array,passenger)
	}
	if(ticket.check()) {
		ticket.printTicket()
		ticket.printTicketFile()
	} else {
		fmt.Printf("Sorry No tickets available for %s to %s :(\n",ticket.src,ticket.dest)
		fmt.Printf("==============================\n\n")
	}
}

