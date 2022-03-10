package main

import (
	"shopping-app/user/db"
	"shopping-app/user/domain"
)

func main() {

	userDb := db.NewUserRepositoryDB([]db.User{})

	userSvc := domain.NewUserService(userDb)

	user := domain.NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

	userSvc.Register(*user)

}
