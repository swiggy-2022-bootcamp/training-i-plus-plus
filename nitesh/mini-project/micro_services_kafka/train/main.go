package main

import (
	"trainService/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.TrainRouter(router)

	router.Run(":8081")
}
