package tests

import(
	"testing"
	"os"
	"strings"
	"healthcareApp/model"
	"healthcareApp/service"
)

type Doctor model.Doctor
type GeneralUser model.GeneralUser
type Patient model.Patient

func TestRegisterDoctor(t *testing.T){
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
	obj := service.Doctor(doctor)
	service.RegisterDoctor(obj)

	dat, err := os.ReadFile("../data/doctor.json")
    if err != nil {
		t.Errorf("File creation failed. Data not stored into the file.")
	}
    data := string(dat)
	if !strings.Contains(data,"rakeshadani@gmail.com") {
		t.Errorf("Doctor profile creation failed.")
	}
}

func TestRegisterGeneralUser(t *testing.T){

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

	dat, err := os.ReadFile("../data/doctor.json")
    if err != nil {
		t.Errorf("File creation failed. Data not stored into the file.")
	}
    data := string(dat)
	if !strings.Contains(data,"rakeshadani@gmail.com") {
		t.Errorf("Doctor profile creation failed.")
	}
}

func TestRegisterPatient(t *testing.T){

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
	
	dat, err := os.ReadFile("../data/doctor.json")
    if err != nil {
		t.Errorf("File creation failed. Data not stored into the file.")
	}
    data := string(dat)
	if !strings.Contains(data,"rakeshadani@gmail.com") {
		t.Errorf("Doctor profile creation failed.")
	}
}