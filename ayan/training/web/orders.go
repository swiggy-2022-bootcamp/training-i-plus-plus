package main

import (
	"encoding/json"
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
	ItemId      string `json:"itemId"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

var orders []Order
var prevOrderId = 0

func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	prevOrderId++
	order.OrderID = strconv.Itoa(prevOrderId)
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
	inputOrderId := params["orderId"]
	for _, order := range orders {
		if order.OrderID == inputOrderId {
			json.NewEncoder(w).Encode(order)
			return
		}
	}

	json.NewEncoder(w).Encode(orders)
}

func updateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderId := params["orderId"]
	for i, order := range orders {
		if order.OrderID == inputOrderId {
			var updatedOrder Order
			json.NewDecoder(r.Body).Decode(&updatedOrder)
			updatedOrder.OrderID = inputOrderId
			updatedOrder.OrderedAt = time.Now()
			orders = append(orders[:i], orders[i+1:]...)
			orders = append(orders, updatedOrder)
			json.NewEncoder(w).Encode(updatedOrder)
			return
		}
	}

	json.NewEncoder(w).Encode(orders)
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderId := params["orderId"]
	for i, order := range orders {
		if order.OrderID == inputOrderId {
			orders = append(orders[:i], orders[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}

func main() {

	router := mux.NewRouter()
	router.handleFunc("/orders", createOrder).Method("POST")
	router.handleFunc("/orders", getOrders).Method("GET")
	router.handleFunc("/orders/{orderId}", updateOrder).Method("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}
