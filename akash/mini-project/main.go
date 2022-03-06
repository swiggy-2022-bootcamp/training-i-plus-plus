package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	model "sample.akash.com/model"
	user "sample.akash.com/user"
)

func main() {
	fmt.Println("Hello ", &model.User{"Ash", "Lambert", "ash.lambert@swiggy.com", "12345"})
	user.Login("abc", "def")

	r := gin.Default()
	fmt.Println(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello-world",
		})
	})
	r.Run()
}
