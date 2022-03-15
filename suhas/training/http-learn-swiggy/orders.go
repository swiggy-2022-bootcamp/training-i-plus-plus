package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Order struct {
	OrderID      string    `json:"orderId"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items"`
}
type Item struct {
	ItemID      string `json:"itemId"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

var orders []Order
var prevOrderID = 0

func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	//var item[] Item
	fmt.Println(r.Body)
	json.NewDecoder(r.Body).Decode(&order)
	//item := order.Items
	fmt.Println(order)
	prevOrderID++
	order.OrderID = strconv.Itoa(prevOrderID)
	order.OrderedAt = time.Now()
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)

}

func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for _, order := range orders {
		if order.OrderID == inputOrderID {
			json.NewEncoder(w).Encode(order)
			return
		}

	}
}

func updateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	var updateOrder Order
	json.NewDecoder(r.Body).Decode(&updateOrder)

	for i, order := range orders {
		if order.OrderID == inputOrderID {
			order.CustomerName = updateOrder.CustomerName
			orders[i] = order
			//orders = append(orders, orders...)
		}
	}
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for i, order := range orders {
		if order.OrderID == inputOrderID {
			orders = append(orders[:i], orders[i+1:]...)
		}
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/orders", createOrder).Methods("POST")
	router.HandleFunc("/orders", getOrders).Methods("GET")
	router.HandleFunc("/orders/{orderId}", getOrder).Methods("GET")
	router.HandleFunc("/orders/{orderId}", updateOrder).Methods("PUT")
	router.HandleFunc("/orders/{orderId}", DeleteOrder).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))

}
