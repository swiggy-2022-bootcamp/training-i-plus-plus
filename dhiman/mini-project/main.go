package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(rw).Encode(map[string]string{"data": "Hello from Mux & mongoDB"})
		if err != nil {
			return
		}
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
