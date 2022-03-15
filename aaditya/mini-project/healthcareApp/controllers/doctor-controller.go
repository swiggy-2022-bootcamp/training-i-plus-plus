package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"healthcareApp/model"
	"healthcareApp/service"
)

type Doctor model.Doctor

func RegisterDoctor(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var doctor Doctor
	//var item[] Item
	//fmt.Println(r.Body)
	json.NewDecoder(r.Body).Decode(&doctor)
	//item := order.Items
	fmt.Println(doctor)
	obj := service.Doctor(doctor)
	service.RegisterDoctor(obj)
	json.NewEncoder(w).Encode(doctor)
}