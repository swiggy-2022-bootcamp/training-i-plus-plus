package service

import (
	"encoding/json"
	"fmt"
	"healthcareApp/model"
	"log"
	"os"
)

func init() {
    file, err := os.Create("data/patient.json")
     
    if err != nil {
        log.Fatalf("failed creating file: %s", err)
    }

	file.Close()

	file, err = os.Create("data/generalUser.json")
     
    if err != nil {
        log.Fatalf("failed creating file: %s", err)
    }

	file.Close()

	file, err = os.Create("data/doctor.json")
     
    if err != nil {
        log.Fatalf("failed creating file: %s", err)
    }

	file.Close()
}

type GeneralUser model.GeneralUser
type Doctor model.Doctor
type Patient model.Patient

type FileWriter interface {
	WriteDataToFile()
}

type FileReader interface {
	ReadDataFromFile()
}


func (generalUser GeneralUser) WriteDataToFile(){
	file,err := os.OpenFile("data/generalUser.json", os.O_APPEND| os.O_CREATE | os.O_WRONLY, 0644)
	handleErr(err)
	defer file.Close()
	data, _ := json.Marshal(generalUser)
	_,err = file.WriteString(string(data))
	handleErr(err)
	fmt.Println("Data written to file successfully.")
	
}


func (doctor Doctor) WriteDataToFile(){
	file,err := os.OpenFile("data/doctor.json", os.O_WRONLY|os.O_APPEND, 0644)
	handleErr(err)
	defer file.Close()
	data, _ := json.Marshal(doctor)
	_,err = file.WriteString(string(data))
	handleErr(err)
	fmt.Println("Data written to file successfully.")
	
}

func (patient Patient) WriteDataToFile(){
	file,err := os.OpenFile("data/patient.json", os.O_WRONLY|os.O_APPEND, 0644)
	handleErr(err)
	defer file.Close()
	data, _ := json.Marshal(patient)
	_,err = file.WriteString(string(data))
	handleErr(err)
	fmt.Println("Data written to file successfully.")
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic("Something went wrong")
	}
}