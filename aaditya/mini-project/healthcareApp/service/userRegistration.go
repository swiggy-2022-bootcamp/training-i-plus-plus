package service

import (
	"fmt"
	"healthcareApp/model"
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

func RegisterDoctor(doctor Doctor){
	doctor.WriteDataToFile()
	doctor.printUserDetails()		
}

func RegisterGeneralUser(generalUser GeneralUser){
	generalUser.WriteDataToFile()
	generalUser.printUserDetails()
}

func RegisterPatient(patient Patient){
	patient.WriteDataToFile()
	patient.printUserDetails()		
}

