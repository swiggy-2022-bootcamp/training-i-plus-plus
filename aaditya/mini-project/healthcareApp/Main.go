package main

import (
	"fmt"
	"healthcareApp/model"
	"healthcareApp/service"
	//"healthcareApp/db"
)

type GeneralUser model.GeneralUser
type Doctor model.Doctor
type Patient model.Patient

func main(){

	// cfg := db.CreateLocalClient()
	// 	fmt.Println(cfg)

	//user registration
	role:="patient"
	if role == "doctor"{
		doctor:= Doctor{
			Id : "1",
			Category : "neurologist",
			Yoe : 12,
			MedicalLicenseLink:  "amazon.s3.com/id",
			User: model.User{
				Id: "1",
				Name : "Dr. Chintan Agarwal",
				EmailId: "chintanagarwal@gmail.com",
				Age : 33,
				Address: "Vapi",
			},
		}
		obj := service.Doctor(doctor)
		service.RegisterDoctor(obj)
	}else if role == "generalUser"{
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
		service.RegisterGeneralUser(obj)
	}else if role == "patient"{
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
		obj := service.Patient(patient)
		service.RegisterPatient(obj)
	}else{
		panic("The given role does not exists.")
	}

	//Add slots for a particular doctor
	service.OpenSlotsForAppointments("Rakesh","01-01-2022 15:00:00",500,false)
	service.OpenSlotsForAppointments("Rakesh","01-01-2022 16:00:00",500,false)
	service.OpenSlotsForAppointments("Anant","01-01-2022 15:00:00",500,false)
	service.OpenSlotsForAppointments("Chintan","01-01-2022 15:00:00",500,false)
	service.OpenSlotsForAppointments("Shloka","01-01-2022 15:00:00",500,false)

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


