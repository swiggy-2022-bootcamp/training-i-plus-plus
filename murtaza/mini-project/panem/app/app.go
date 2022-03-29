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

	authHandler := AuthHandler{
		authService: domain.NewAuthService(userMongoRepository),
	}

	r := Routes{
		router: gin.Default(),
	}

	v1 := r.router.Group("/v1")

	users := v1.Group("/users")

	users.Use(authHandler.authMiddleware)
	users.GET("/", userHandler.getAllUsers)
	users.GET("/:userId", userHandler.getUserByUserId)
	users.DELETE("/:userId", userHandler.deleteUser)
	users.PUT("/:userId", userHandler.updateUser)

	login := v1.Group("/login")
	login.POST("/", authHandler.handleLogin)

	signup := v1.Group("/signup")
	signup.POST("/signup", userHandler.createUser)

	auth := v1.Group("/auth")
	auth.GET("/", authHandler.isTokenValid)

	r.router.Run(":8089")
}
