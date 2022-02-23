package customer

import (
	// "BookingApp/train"
	// "BookingApp/ticket"
	"fmt"
	"strconv"

	"encoding/json"
	scribble "github.com/nanobox-io/golang-scribble"
)

func AddCustomer(Cust Customer) bool {
	var customer_dir string = "./customer/customer_json"
	customer_db, err := scribble.New(customer_dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}
	if err := customer_db.Write("customer", strconv.Itoa(Cust.Customer_id), Cust); err != nil {
		//log error
		fmt.Println(err)
		return false
	}
	return true
}

func GetCustomer(id int) (Customer,error) {
	customer_obj := Customer{}
	var customer_dir string = "./customer/customer_json"
	customer_db, err := scribble.New(customer_dir, nil)
	if err != nil {
		fmt.Println("Error", err)
		return customer_obj,err
	}

	if err := customer_db.Read("customer", strconv.Itoa(id), &customer_obj); err != nil {
		//log error
		fmt.Println(err)
		return customer_obj,err
	}
	return customer_obj,nil
}

func GetCustomers() []Customer {
	var customer_dir string = "./customer/customer_json"
	customer_db, err := scribble.New(customer_dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}
	var customer_objs []string
	customer_objs, err = customer_db.ReadAll("customer")
	if err != nil {
		//log error
		fmt.Println(err)
	}
	var customers = []Customer{}
	for _, f := range customer_objs {
	customerFound := Customer{}
	if err := json.Unmarshal([]byte(f), &customerFound); err != nil {
		fmt.Println("Error", err)
	}
	customers = append(customers, customerFound)
	}
	return customers
}

func DeleteCustomer(cust Customer) {
	customer_array := GetCustomers()
	for i,v := range(customer_array) {
		if v.Customer_id == cust.Customer_id && v.Firstname == cust.Firstname && v.Lastname == cust.Lastname {
			customer_array = append(customer_array[:i], customer_array[i+1:]...)
			for _,cust := range(customer_array) {
				AddCustomer(cust)
			}
		}
	}
}

func (cust Customer) UpdateCustomers(newcustomer Customer) bool {
	DeleteCustomer(cust)
	AddCustomer(newcustomer)
	return true
}
