package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"healthcareApp/model"
	"healthcareApp/service"
)


type GeneralUser model.GeneralUser

func RegisterGeneralUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var generalUser GeneralUser
	//var item[] Item
	//fmt.Println(r.Body)
	json.NewDecoder(r.Body).Decode(&generalUser)
	//item := order.Items
	fmt.Println(generalUser)
	obj := service.GeneralUser(generalUser)
	service.RegisterGeneralUser(obj)
	json.NewEncoder(w).Encode(generalUser)
}