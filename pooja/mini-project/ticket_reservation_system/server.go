package main

import (
	"log"
	"net/http"
	user_controller "ticket_reservation_system/controller"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", user_controller.AddUser).Methods("POST")
	router.HandleFunc("/user", user_controller.GetUsers).Methods("GET")
	router.HandleFunc("/user/{user_id}", user_controller.GetUserById).Methods("GET")
	router.HandleFunc("/user/{user_id}", user_controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{user_id}", user_controller.DeleteUserById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5001", router))
}
