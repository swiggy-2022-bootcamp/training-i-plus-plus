package main

import (
	"gin-mongo-api/config" //add this
	"gin-mongo-api/routes" 
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    
    //run database
    config.ConnectDB()

    //routes
    routes.UserRoute(router) //add this
    routes.AdminRoute(router) //add this

    router.Run("localhost:6000")
}