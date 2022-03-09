package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func index(w http.ResponseWriter, req *http.Request)  {
    fmt.Fprintf(w, "Welcome to MyNotePaper!")
}

func posts(w http.ResponseWriter, req *http.Request)  {
    vars := mux.Vars(req)
    post_id := vars["id"]

    fmt.Fprintf(w, "The requested post ID: %s\n", post_id)
}

func main() {
    r := mux.NewRouter()
    // routes
    r.HandleFunc("/", index)
    r.HandleFunc("/posts/{id}", posts)

    // listen port 80
    http.ListenAndServe(":8080", r)
}