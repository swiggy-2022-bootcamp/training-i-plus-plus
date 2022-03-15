package routes

import (
    "gin-mongo-api/controllers"
    "github.com/gin-gonic/gin"
)

func AdminRoute(router *gin.Engine) {
    router.POST("/admin", controllers.CreateAdmin())
    router.GET("/admin/:adminId", controllers.GetAAdmin())
    router.PUT("/admin/:adminId", controllers.EditAAdmin())
    router.DELETE("/admin/:adminId", controllers.DeleteAAdmin())
    router.GET("/admins", controllers.GetAllAdmins())
}