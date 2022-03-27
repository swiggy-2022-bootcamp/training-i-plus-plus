package main

import (
	"fmt"
	"srctc/config"
	"srctc/routes"

	"github.com/gin-gonic/gin"
)

func UserRouter() *gin.Engine {
	fmt.Println("user check")
	router := gin.Default()
	routes.UserRoute(router)
	return router
}

func AdminRouter() *gin.Engine {
	fmt.Println("admin check")
	router := gin.Default()
	routes.AdminRoute(router)
	return router
}

func AuthRouter() *gin.Engine {
	router := gin.Default()
	routes.AuthRoute(router)
	return router
}

func main() {

	config.ConnectDB()

	fmt.Println("Type  1. User  2. Admin  3. Auth: ")
	var choice int
	fmt.Scanln(&choice)

	if choice == 1 {
		userroute := UserRouter()
		userroute.Run("localhost:6000")
	}
	if choice == 2 {
		adminroute := AdminRouter()
		adminroute.Run("localhost:7000")
	}
	if choice == 3 {
		authroute := AuthRouter()
		authroute.Run("localhost:8000")
	}

}
