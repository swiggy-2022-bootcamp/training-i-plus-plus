package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"mini-project/internal/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", handlers.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users", handlers.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", handlers.GetUserByID).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", handlers.UpdateUserByID).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", handlers.DeleteUserByID).Methods(http.MethodDelete)

  log.Println("API is running in Port 4000!")
  http.ListenAndServe(":4000", router)
}