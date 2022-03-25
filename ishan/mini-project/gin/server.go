package main

import (
	"fmt"
	JWTManager "swiggy/gin/lib/helpers"
	db "swiggy/gin/lib/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	client, ctx, cancel := db.ConnectDB()
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
		fmt.Println("MongoDB Connection Closed")
	}()

	JWTManager.NewJWTManager("Ishan", time.Hour*50)

	router := gin.Default()
	ApplyRoutes(router)

	router.Run("localhost:8080")
}
