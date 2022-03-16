package routes

import (
	"healthcareApp/controllers"
)


func GeneralUserRoutes(){	
	Router.HandleFunc("/generalUserRegistration",controllers.RegisterGeneralUser).Methods("POST")
}

