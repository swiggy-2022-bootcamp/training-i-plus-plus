package router

import (
	"medo-healthcare-app/cmd/controller"
	"medo-healthcare-app/pkg/authentication"

	"github.com/gorilla/mux"
)

//Router ...
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.ServeHome).Methods("GET")
	router.HandleFunc("/login", authentication.ValidateLogin).Methods("GET")
	router.HandleFunc("/createUser", controller.CreateUser).Methods("POST")
	router.HandleFunc("/getAllUsers", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/getUser", controller.GetOneUser).Methods("GET")
	router.HandleFunc("/updateUser", controller.UpdateOneUser).Methods("PUT")
	router.HandleFunc("/wipeAllUsersData", controller.DeleteAllUsers).Methods("DELETE")
	//router.HandleFunc("/getUsername", controller.GetUsernameFromTokenController).Methods("GET")
	return router
}
