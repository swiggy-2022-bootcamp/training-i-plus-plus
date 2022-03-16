package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"healthcareApp/model"
	"healthcareApp/service"
)

type Patient model.Patient

func RegisterPatient(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var patient Patient
	//var item[] Item
	//fmt.Println(r.Body)
	json.NewDecoder(r.Body).Decode(&patient)
	//item := order.Items
	fmt.Println(patient)
	obj := service.Patient(patient)
	service.RegisterPatient(obj)
	json.NewEncoder(w).Encode(patient)
}