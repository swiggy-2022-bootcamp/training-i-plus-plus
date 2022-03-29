package main

import (
	"context"
	kf "swiggy/gin/goKafka/consumer"
	db "swiggy/gin/lib/utils"
	router "swiggy/gin/router"
)

func init() {
	db.ConnectDB()
}
func main() {

	ctx := context.Background()

	kf.Consume(ctx)

	router.ApplyRoutes().Run("localhost:8081")
}
