package user

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	// "sample.akash.com/model"
)

func Login(c *gin.Context) {
	fmt.Println("Login ", c.Request.Body)

	var data = jsonparser.Get(c.Request.Body, "Email", "Password")

	fmt.Println(data)
}
