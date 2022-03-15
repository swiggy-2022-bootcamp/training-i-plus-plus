package routes

import (
	"healthcareApp/controllers"
)

func DoctorRoutes(){	
	Router.HandleFunc("/doctorRegistration",controllers.RegisterDoctor).Methods("POST")
}

