package main

import (
	"userService/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.UserRouter(router)

	router.Run(":8080")
}
