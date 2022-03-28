package main

import (
	"aman-swiggy-mini-project/models"
	"aman-swiggy-mini-project/routes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting Project")
	r := gin.New()
	r.Use(gin.Logger())
	routes.UserRoutes(r)
	routes.SellerRoutes(r)
	routes.ProductRoutes(r)
	routes.CartRoutes(r)
	r.POST("/", test)
	r.Run()
}

func test(c *gin.Context) {
	body := c.Request.Body
	decoder := json.NewDecoder(body)
	var user1 models.User
	err := decoder.Decode(&user1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user1)
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"HI":   value,
		"body": user1,
	})
}
