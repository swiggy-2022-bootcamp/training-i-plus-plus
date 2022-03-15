package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Order struct {
	OrderId      string    `json:"orderId"`
	CustomerName string    `json:"customerName"`
	OrderedTime  time.Time `json:"orderAt"`
	Items        []Item    `json:"item"`
}

type Item struct {
	ItemId      string `json:"itemId"`
	Description string `json:"desc"`
	Quantity    int    `json:"quantity"`
}

var orders []Order

func GetOrdersHandlers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func CreateOrderHandlers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var order Order
	_ = json.NewDecoder(r.Body).Decode(&order)
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetOrdersHandlers).Methods("GET")
	router.HandleFunc("/", CreateOrderHandlers).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
