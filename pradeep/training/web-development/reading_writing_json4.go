// Unmarshalling json to structs.
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type helloWorldRequest struct {
	Name string
}

type helloWorldResponse struct {
	Message string
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	response := helloWorldResponse{Message: "Hello " + request.Name}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func main() {
	port := 8080
	http.HandleFunc("/hello", helloWorldHandler)
	log.Printf("Listening on port %v", port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
