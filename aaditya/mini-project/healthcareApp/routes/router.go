package routes

import (
	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

func InitializeRouter(){
	DoctorRoutes()
	PatientRoutes()
	GeneralUserRoutes()
}