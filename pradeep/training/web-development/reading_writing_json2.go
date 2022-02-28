package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string
}

// Is there any better way to send our data to the output stream without marshalling to a temporary object before we return it? Yes.
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello World"}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}

func main() {
	port := 8080
	http.HandleFunc("/hello", helloWorldHandler)
	log.Printf("Listening on port %v", port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
