package controller

import (
	"encoding/json"
	"medo-healthcare-app/cmd/model"
	"medo-healthcare-app/pkg/database"
	"net/http"
)

//ServeHome ..
func ServeHome(w http.ResponseWriter, r *http.Request) {
	homePage := "./web/assets/html/homePage.html"
	http.ServeFile(w, r, homePage)
}

//CreateUser ..
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var user model.CoreUserData
	_ = json.NewDecoder(r.Body).Decode(&user)
	database.InsertOne(user)
	json.NewEncoder(w).Encode(user)
}

//GetAllUsers ...
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	allUsersData := database.Find()
	json.NewEncoder(w).Encode(allUsersData)
}

//GetOneUser ...
func GetOneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.CoreUserData
	_ = json.NewDecoder(r.Body).Decode(&user)
	result := database.FindOne(user.Email)
	json.NewEncoder(w).Encode(result)
}
