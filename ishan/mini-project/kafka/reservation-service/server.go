package main

import (
	db "swiggy/gin/lib/utils"

	router "swiggy/gin/router"
)

func main() {
	db.ConnectDB()

	router.ApplyRoutes().Run("localhost:8080")
}
