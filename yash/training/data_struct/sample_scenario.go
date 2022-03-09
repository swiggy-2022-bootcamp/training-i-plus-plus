package main

import "fmt"

type fees interface {
	total_fees() float64
	doctor_fees(total_fees float64) float64
}
type doctor struct {
	name      string
	age       int
	fees      int
	exp       int
	expertise string
}
type user struct {
	name     string
	age      int
	program  string
	goal     string
	doc      doctor
	facility []string
}

func (u user) total_fees() float64 {
	var total_fees float64
	total_fees += float64(u.doc.fees)
	if u.program == "medical" {
		total_fees += 100
	} else if u.program == "dental" {
		total_fees += 200
	} else if u.program == "general" {
		total_fees += 50
	}
	total_fees += float64(u.age) * 0.8
	total_fees += float64(u.doc.exp) * 0.5
	for i := 0; i < len(u.facility); i++ {
		if u.facility[i] == "hospital" {
			total_fees += 100
		} else if u.facility[i] == "test" {
			total_fees += 50
		} else if u.facility[i] == "pharmacy" {
			total_fees += 30
		}

	}
	return total_fees
}

func (u user) doctor_fees(total_fees float64) float64 {
	var fee float64
	fee = total_fees * 0.1
	return fee
}
func main() {
	var u1, u2 fees
	d1 := doctor{"yash", 25, 1000, 5, "general"}
	d2 := doctor{"raj", 20, 2000, 15, "general"}
	u1 = user{"patient1", 25, "medical", "weight gain", d1, []string{"hospital", "test"}}
	u2 = user{"patient2", 30, "dental", "weight loss", d2, []string{"hospital", "pharamcy"}}

	fmt.Println("The Fees for the first patient is:", u1.total_fees())
	fmt.Println("The Fees for the second patient is:", u2.total_fees())
	fmt.Println("The first Doctor Fees is:", u1.doctor_fees(u1.total_fees()))
	fmt.Println("The second Doctor Fees is:", u2.doctor_fees(u2.total_fees()))
}
