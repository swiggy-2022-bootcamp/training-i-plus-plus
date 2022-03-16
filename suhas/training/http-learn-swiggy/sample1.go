package main

import (
	"fmt"
	"net/http"
)

func hi(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	http.HandleFunc("/hi", hi)
	http.ListenAndServe(":5000", nil)
}
