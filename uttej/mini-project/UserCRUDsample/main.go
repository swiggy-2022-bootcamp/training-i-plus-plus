package main

import (
	"golang-app/controllers"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func main() {

	uc := controllers.NewUserController(getSession())

	router := gin.Default()
	router.GET("/users", uc.UsersList)
	router.GET("/users/:id", uc.GetUser)
	router.DELETE("/users/:id", uc.RemoveUser)
	router.POST("/users", uc.CreateUser)
	router.PUT("/users/:id", uc.UpdateUser)

	router.Run(":4000")
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return s
}
