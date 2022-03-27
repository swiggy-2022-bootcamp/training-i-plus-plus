package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("All users printed")
		c.JSON(200, "allUsers")
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("User with id printed")
		c.JSON(200, "allUsers")
	}
}

func SignUp(userData string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(userData)
		c.JSON(200, "allUsers")
	}
}

func Login(userCred string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(userCred)
		c.JSON(200, "allUsers")
	}
}
