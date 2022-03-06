package main

import (
	"fmt"
	"healthcareApp/model"
	"healthcareApp/service"
)

type GeneralUser model.GeneralUser
type Doctor model.Doctor
type Patient model.Patient


type UserDetails interface{
	printUserDetails()
}

func (generalUser GeneralUser) printUserDetails(){
	fmt.Println("User registed successfully with the following details. ", generalUser)	
}

func (patient Patient) printUserDetails(){
	fmt.Println("Patient registed successfully with the following details. ", patient)
}

func (doctor Doctor) printUserDetails(){
	fmt.Println("Doctor registed successfully with the following details. ", doctor)
}

func main(){

	//user registration
	role:="generalUser"
	if role == "doctor"{
		registerDoctor()
	}else if role == "generalUser"{
		registerGeneralUser()
	}else if role == "patient"{
		registerPatient()
	}else{
		panic("The given role does not exists.")
	}

	// Add slots for a particular doctor
	openSlotsForAppointments("Rakesh","01-01-2022 15:00:00",500,false)
	openSlotsForAppointments("Rakesh","01-01-2022 16:00:00",500,false)
	openSlotsForAppointments("Anant","01-01-2022 15:00:00",500,false)
	openSlotsForAppointments("Chintan","01-01-2022 15:00:00",500,false)
	openSlotsForAppointments("Shloka","01-01-2022 15:00:00",500,false)

	//getAllSlots
	service.GetAllSlots()

	//book appointment by doctor name
	service.BookAppointmentsByDoctorName("Rakesh")

	//getAllOpenSlots
	service.GetAllOpenSlots()

	//book appointment by availability
	service.BookAppointmentsByOpenSlots()

	//getAllOpenSlots
	service.GetAllOpenSlots()

	//Waiting for all go routines to exit.
	service.Wg.Wait()
	fmt.Println("Program ended")
}

func registerDoctor(){
	doctor:= Doctor{
		Id : "1",
		Category : "surgeon",
		Yoe : 12,
		MedicalLicenseLink:  "amazon.s3.com/id",
		User: model.User{
			Id: "1",
			Name : "Dr. Rakesh Adani",
			EmailId: "rakeshadani@gmail.com",
			Age : 43,
			Address: "Mumbai",
		},
	}
	doctor.printUserDetails()		
}

func registerGeneralUser(){
	generalUser := GeneralUser{
		IsPatient: false,
		PreviousDiseases: "typhoid",
		Id: "1",
		User: model.User{
			Id: "1",
			Name : "Stokes ",
			EmailId: "johnsteve@gmail.com",
			Age : 27,
			Address: "England",
		},
	}
	obj := service.GeneralUser(generalUser)
	obj.WriteDataToFile()
	generalUser.printUserDetails()
}

func registerPatient(){
	patient:= Patient{
		Id: "1",
		DoctorAssignedId: "1",
		RoomAllocated: "702E",
		IsDischarged: false,
		User: model.User{
			Id: "1",
			Name : "Rakesh Adani",
			EmailId: "rakeshadani@gmail.com",
			Age : 27,
			Address: "Mumbai",
		},
	}
	patient.printUserDetails()		
}

func openSlotsForAppointments(doctorName , slot string, fees int, occupied bool){
	
	appointments := service.Appointments{
		DoctorName : doctorName,
		Slot	   : slot,
		Fees	   : fees,
		Occupied   : occupied,
	}
	newAppointments := []service.Appointments{}
	newAppointments = append(newAppointments, appointments)
	//fmt.Println(newAppointments)
	service.AddSlots(doctorName,newAppointments)
}

