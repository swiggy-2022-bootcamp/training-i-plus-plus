package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"medo-healthcare-app/cmd/router"
	"medo-healthcare-app/docs"
	"net/http"
	"os"
)

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Medo - The Healthcare Application"
}

//Start ..
func Start() {
	fmt.Println("Hey, there !")
	fmt.Println("Medo - Your Onestop Healthcare Point !")

	//Custom Logger - Logs actions to 'shippingAddressService.logger' file
	file, err := os.OpenFile("./logs/medoAppLogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}
	configureSwaggerDoc()
	route := router.Router()
	log.Fatal(http.ListenAndServe(":9001", route))
	//authentication.GetUsernameFromToken(w http.ResponseWriter, r http.Request)
}
