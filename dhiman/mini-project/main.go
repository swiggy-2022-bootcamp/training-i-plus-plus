package main

import (
	"net/http"

	"github.com/dhi13man/healthcare-app/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(
			http.StatusOK,
			models.NewDisease(
				"Covid-19",
				[]string{},
				[]string{"Fever", "Cough", "Shortness of breath"},
			),
		)
	})

	router.Run()
}
