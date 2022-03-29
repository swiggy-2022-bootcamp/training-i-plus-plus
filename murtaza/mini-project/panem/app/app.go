package app

import (
	"panem/domain"
	"panem/infra"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	router *gin.Engine
}

func Start() {

	userMongoRepository := infra.NewUserMongoRepository()

	userHandler := UserHandler{
		userService: domain.NewUserService(userMongoRepository),
	}

	r := Routes{
		router: gin.Default(),
	}

	v1 := r.router.Group("/v1")

	users := v1.Group("/users")

	users.GET("/", userHandler.getAllUsers)
	users.GET("/:userId", userHandler.getUserByUserId)
	users.POST("/", userHandler.createUser)
	users.DELETE("/:userId", userHandler.deleteUser)
	users.PUT("/:userId", userHandler.updateUser)

	r.router.Run(":8089")
}
