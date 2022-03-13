package infra

import "testing"

var mockUserRepository = NewUserRepository()

func TestShouldGetAllUsers(t *testing.T) {
	actualResponse := userRepository.GetAllUsers()
	expectedResponse := 
}
