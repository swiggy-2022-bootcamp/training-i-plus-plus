package controller

import (
	"encoding/json"
	"medo-healthcare-app/pkg/authentication"
	"medo-healthcare-app/pkg/database"
	"medo-healthcare-app/pkg/logger"
	"net/http"
)

//GetAllDoctors ...
func GetAllDoctors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	if authentication.AuthenticateLogin(w, r) {
		if (GetUserType(w, r)) == "patient" || GetUserType(w, r) == "masteradmin" || GetUserType(w, r) == "admin" {
			allDoctorsData := database.DocFind()
			logger.Info("Doctors Data - FETCHED SUCCESSFULLY !")
			json.NewEncoder(w).Encode(allDoctorsData)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "Unauthorized Access.\nPlease Try Logging Again with Administrator Access :)", http.StatusUnauthorized)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "Unauthorized Access.\nPlease LOGIN to access this URL ! :)", http.StatusUnauthorized)
	}
}
