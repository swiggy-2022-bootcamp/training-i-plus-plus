package main

import (
	"fmt"
)

type User struct{
	id			string
	name 		string
	emailId 	string
	age 		int
	address		string 
}

type Doctor struct{
	id					string
	User
	category 			string
	yoe 	 			float64
	medicalLicenseLink	string
}

type GeneralUser struct{
	id 					string
	User
	previousDiseases	string
	isPatient			bool

}

type Patient struct {
	id 					string
	User
	doctorAssignedId	string
	isDischarged		bool
	roomAllocated		string
}

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
	role:="doctor"
	if role == "doctor"{
		registerDoctor()
	}else if role == "generalUser"{
		registerGeneralUser()
	}else if role == "patient"{
		registerPatient()
	}else{
		panic("The given role does not exists.")
	}
}

func registerDoctor(){

	doctor:= Doctor{
		id : "1",
		category : "surgeon",
		yoe : 12,
		medicalLicenseLink:  "amazon.s3.com/id",
		User: User{
			id: "1",
			name : "Dr. Rakesh Adani",
			emailId: "rakeshadani@gmail.com",
			age : 43,
			address: "Mumbai",
		},
	}

	doctor.printUserDetails()
		
}

func registerGeneralUser(){
	generalUser := GeneralUser{
		isPatient: false,
		previousDiseases: "asthma",
		id: "1",
		User: User{
			id: "1",
			name : "Rakesh Adani",
			emailId: "rakeshadani@gmail.com",
			age : 27,
			address: "Mumbai",
		},
	}
	
	generalUser.printUserDetails()
}

func registerPatient(){
	patient:= Patient{
		id: "1",
		doctorAssignedId: "1",
		roomAllocated: "702E",
		isDischarged: false,
		User: User{
			id: "1",
			name : "Rakesh Adani",
			emailId: "rakeshadani@gmail.com",
			age : 27,
			address: "Mumbai",
		},
	}
	
	patient.printUserDetails()
		
}