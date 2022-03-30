package controllers

import (
	"tejas/services"

	"github.com/gin-gonic/gin"
)

func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := services.HealthCheck()
		c.JSON(200, response)
	}
}

func DeepHealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := services.DeepHealthCheck()
		c.JSON(200, response)
	}
}
