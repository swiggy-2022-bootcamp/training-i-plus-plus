package service

import (
	"fmt"
	"sync"
)

var Wg sync.WaitGroup

type Appointments struct {
	DoctorName		string
	Slot			string
	Fees			int
	Occupied		bool
}

type AppointmentsAvailable []Appointments

var AppointmentsMappedByDoctorName map[string]AppointmentsAvailable

func AddSlots(doctorName string, openedSlots AppointmentsAvailable){
	if AppointmentsMappedByDoctorName == nil {
		AppointmentsMappedByDoctorName = make(map[string]AppointmentsAvailable)
	}
	slots, doesExists := AppointmentsMappedByDoctorName[doctorName]

	if !doesExists {
		AppointmentsMappedByDoctorName[doctorName] = openedSlots
	}else{
		slots = append(slots, openedSlots...)
		AppointmentsMappedByDoctorName[doctorName] = slots
	}
	
	//fmt.Println(AppointmentsByDoctor)
}

func GetAllSlots(){
	for doctorName,slotsByDoctor := range AppointmentsMappedByDoctorName {
		for _,slots := range slotsByDoctor{
			slotAvailable:= ""
			if slots.Occupied {
				slotAvailable = "No"
			}else{
				slotAvailable = "Yes"
			}
			fmt.Printf(" Doctor Name - %s\n Available slot - %s\n Fees - %d\n Occupied - %s\n", doctorName, slots.Slot, slots.Fees, slotAvailable )
		}
		
	}
}

func GetAllOpenSlots() {
	for doctorName,slotsByDoctor := range AppointmentsMappedByDoctorName {
		for _,slots := range slotsByDoctor{
			if !slots.Occupied {
				fmt.Printf(" Doctor Name - %s\n Available slot - %s\n Fees - %d\n", doctorName, slots.Slot, slots.Fees )
			}
			
		}
		
	}
}

func BookAppointmentsByDoctorName(doctorName string) {
	slots, doesExists := AppointmentsMappedByDoctorName[doctorName]
	slotAvailable := false
	if doesExists {
		for i,slot := range slots{
			if !slot.Occupied {
				slot.Occupied = true
				Wg.Add(1)
				go sendAppointmentDetailsOnMail(slot);
				slotAvailable = true
				fmt.Println("Slot booked successfully.")
				slots[i] = slot
				return
			}
		}
		
		if !slotAvailable {
			fmt.Println("All the slots for the given doctor are occupied. Please check back later for new slots")
		}
	}else {
		fmt.Println("Doctor doesn't exists with given name")
	}
}

func BookAppointmentsByOpenSlots(){
	slotAvailable := false
	for _,slots := range AppointmentsMappedByDoctorName {
		for i,slot := range slots{
			if !slot.Occupied {
				slot.Occupied = true
				Wg.Add(1)
				go sendAppointmentDetailsOnMail(slot);
				slotAvailable = true
				fmt.Println("Slot booked successfully.")
				slots[i] = slot
				return
			}
			
		}
		
	}
	if !slotAvailable {
		fmt.Println("All the slots for the given doctor are occupied. Please check back later for new slots")
	}
}

func sendAppointmentDetailsOnMail(slot Appointments){
	//send a mail to the generaluser on a separate go routine.
	defer Wg.Done()
	fmt.Println("Your appointment details are as follows : ")
	fmt.Printf(" Doctor Name - %s\n Appointment time - %s\n Fees - %d\n", slot.DoctorName, slot.Slot, slot.Fees)
	fmt.Println("Mail sent successfully.")

}