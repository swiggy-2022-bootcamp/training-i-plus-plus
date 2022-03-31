package main

import (
	"net/http"
	"time"

	"github.com/dhi13man/healthcare-app/users_service/controllers"
	// users_models "github.com/dhi13man/healthcare-app/users_service/models"
	bookkeeping_models "github.com/dhi13man/healthcare-app/bookkeeping_service/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(
			http.StatusOK,
			bookkeeping_models.NewDisease(
				"Covid-19",
				[]string{},
				[]string{"Fever", "Cough", "Shortness of breath"},
			),
		)
	})

	router.POST("/users/", controllers.CreateClient)

	go router.Run("localhost/users_service:8081")
	time.Sleep(100 * time.Second)
}
