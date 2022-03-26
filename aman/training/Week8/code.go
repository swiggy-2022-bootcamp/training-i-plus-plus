package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/teachers", teacherHandler)
	sHandler := studentHandler{}
	mux.Handle("/v1/students", sHandler)
	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	s.ListenAndServe()
}

func teacherHandler(res http.ResponseWriter, req *http.Request) {
	data := []byte("v1 of teacher's called")
	res.WriteHeader(200)
	res.Write(data)
}

type studentHandler struct{}

func (h studentHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("v1 of student's called")
	res.WriteHeader(200)
	res.Write(data)
}
