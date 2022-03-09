package test

import (
	"fmt"
	"BookingAppMongo/database"
	"BookingAppMongo/models"
	"testing"
)


func TestAddAdmin(t *testing.T) {
	type testCaseAdd struct {
		a models.Admin
		b error
	}
	testAdd := []testCaseAdd{
		{a:models.Admin{},b:nil},
		{a:models.Admin{Name:"admin501",AdminId:"AD501"},b:nil},
		{a:models.Admin{Name:"admin503",AdminId:"AD504"},b:nil},
		{a:models.Admin{Name:"",AdminId:"AD506"},b:nil},
		{a:models.Admin{Name:"admin505",AdminId:""},b:nil},
	}
	for _,v := range(testAdd) {
		res := database.InsertAdmin(v.a)
		if res != v.b{
			t.Error(res)
			t.Errorf("Expected nil but got error")
		}
	}
}

func TestReadAdmin(t *testing.T) {
	type testCaseRead struct {
		a string
		b error
	}
	testRead := []testCaseRead{
		{a:"",b:nil},
		{a:"AD501",b:nil},
		{a:"AD504",b:nil},
		{a:"AD506",b:nil},
	}
	for _,v := range(testRead) {
		res := database.ReadAdmin(v.a)
		fmt.Println(res)
	}
}

func TestReadAllAdmin(t *testing.T) {
	res := database.ReadAllAdmin()
	fmt.Println(res)
}

func TestUpdateAdmin(t *testing.T) {
	type testCaseUpdate struct {
		id string
		nadmin models.Admin
		e   error
	}
	testUpdate := []testCaseUpdate{
		{id:"AD501",nadmin:models.Admin{Name:"admin5901",AdminId:"AD5901"},e:nil},
		{id:"AD503",nadmin:models.Admin{Name:"admin5903",AdminId:"AD5903"},e:nil},
		{id:"AD506",nadmin:models.Admin{Name:"admin5906",AdminId:"AD5906"},e:nil},
	}
	for _,v := range(testUpdate) {
		res := database.UpdateAdmin(v.id,v.nadmin)
		if res != v.e{
			t.Error(res)
			t.Errorf("Expected nil but got error")
		}
	}
}

func TestDeleteAdmin(t *testing.T) {
	type testCaseDelete struct {
		id string
		e error
	}
	testDelete := []testCaseDelete{
		{id:"AD5901",e:nil},
		{id:"AD5903",e:nil},
	}
	for _,v := range(testDelete) {
		res := database.DeleteAdmin(v.id)
		if res != v.e{
			t.Error(res)
			t.Errorf("Expected nil but got error")
		}
	}
}