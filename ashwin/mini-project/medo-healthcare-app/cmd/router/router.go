package router

import (
	"medo-healthcare-app/cmd/controller"

	"github.com/gorilla/mux"
)

//Router ...
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.ServeHome).Methods("GET")
	router.HandleFunc("/createUser", controller.CreateUser).Methods("POST")
	router.HandleFunc("/allUsers", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/getUser", controller.GetOneUser).Methods("GET")
	return router
}
