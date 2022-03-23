package main

import "fmt"

func hi(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	handler := http.NewServMux()
	handler.handleFunc("/hi", hi)
	http.ListedAndServe(":8080", nil)
}
