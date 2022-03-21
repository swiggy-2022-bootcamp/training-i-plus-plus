package test

import (
	"fmt"
	"BookingAppMongo/database"
	"BookingAppMongo/models"
	"testing"
)


func TestAddCustomer(t *testing.T) {
	type testCaseAdd struct {
		a models.Customer
		b error
	}
	testAdd := []testCaseAdd{
		{a:models.Customer{},b:nil},
		{a:models.Customer{Firstname:"Joe",Lastname:"Toed",CustomerId:"CU501"},b:nil},
		{a:models.Customer{Firstname:"Koe",Lastname:"Uoed",CustomerId:"CU504"},b:nil},
		{a:models.Customer{Firstname:"Loe",Lastname:"Voed",CustomerId:"CU506"},b:nil},
		{a:models.Customer{Firstname:"Moe",Lastname:"Woed",CustomerId:""},b:nil},
	}
	for _,v := range(testAdd) {
		res := database.InsertCustomer(v.a)
		if res != v.b{
			t.Error(res)
			t.Errorf("Expected nil but got error")
		}
	}
}

func TestReadCustomer(t *testing.T) {
	type testCaseRead struct {
		a string
		b error
	}
	testRead := []testCaseRead{
		{a:"",b:nil},
		{a:"CU501",b:nil},
		{a:"CU504",b:nil},
		{a:"CU506",b:nil},
	}
	for _,v := range(testRead) {
		res := database.ReadCustomer(v.a)
		fmt.Println(res)
	}
}

func TestReadAllCustomer(t *testing.T) {
	res := database.ReadAllCustomer()
	fmt.Println(res)
}

func TestUpdateCustomer(t *testing.T) {
	type testCaseUpdate struct {
		id string
		nadmin models.Customer
		e   error
	}
	testUpdate := []testCaseUpdate{
		{id:"CU501",nadmin:models.Customer{Firstname:"Joe",Lastname:"Toed",CustomerId:"CU5901"},e:nil},
		{id:"CD503",nadmin:models.Customer{Firstname:"Koe",Lastname:"Uoed",CustomerId:"CU5907"},e:nil},
		{id:"CD506",nadmin:models.Customer{Firstname:"Loe",Lastname:"Voed",CustomerId:"CU5911"},e:nil},
	}
	for _,v := range(testUpdate) {
		res := database.UpdateCustomer(v.id,v.nadmin)
		if res != v.e{
			t.Error(res)
			t.Errorf("Expected nil but got error")
		}
	}
}

func TestDeleteCustomer(t *testing.T) {
	type testCaseDelete struct {
		id string
		e error
	}
	testDelete := []testCaseDelete{
		{id:"CU5901",e:nil},
		{id:"CU5911",e:nil},
	}
	for _,v := range(testDelete) {
		res := database.DeleteCustomer(v.id)
		if res != v.e{
			t.Error(res)
			t.Errorf("Expected nil but got error")
		}
	}
}