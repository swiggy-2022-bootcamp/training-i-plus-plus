package routes

import (
	"healthcareApp/controllers"
)

func PatientRoutes(){	
	Router.HandleFunc("/patientRegistration",controllers.RegisterPatient).Methods("POST")
}

