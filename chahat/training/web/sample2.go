package web

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)
type Order struct{
	OrderId string `json:"orderId"`
	Items []Item `json:"items"`
	created_at time.Time `json:"created_at"`
	username string `json:"username"`
 
}
type Item struct{
	ItemId 

}
func createOrder(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	prevOrderId++
	order.OrderId=strconv.Itoa(prevOrderId)
	orders=append(orders,order)
	json.NewEncoder(w).Encode(order)

} 
func getOrderById (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	inputId :=params["order_id"]
	for _,order :=range orders {
		if order.OrderId == inputId {
			json.NewEncoder(w).Encode(order)
			return
		}
	}
}
func getOrders (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(orders)
}
var orders []Order
var prevOrderId=0
func main(){
	router := mux.NewRouter()
	router.HandleFunc("/orders",createOrder()).Methods("POST")
	log.Fatal(http.ListenAndServe(":5001",router))
}