package main

import (
	"fmt"
	"mini-project/domain"
	"mini-project/infra"
)

func main() {
	userRepository := infra.NewUserRepository()
	userService := domain.NewUserService(userRepository)

	firstName := "Murtaza"
	lastName := "Sadriwala"
	phone := "9900887766"
	email := "murtaza896@gmail.com"
	username := "murtaza896"
	password := "Pass!23"
	role := domain.Admin

	user, _ := userService.CreateUser(firstName, lastName, phone, email, username, password, role)
	userPersistedEntity, _ := userRepository.FindByEmail(user.Email())
	fmt.Println(userPersistedEntity)
}
