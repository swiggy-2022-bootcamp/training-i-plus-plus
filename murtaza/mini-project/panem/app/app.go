package app

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"panem/docs"
	"panem/domain"
	"panem/infra"
	"panem/utils/logger"
)

type Routes struct {
	router *gin.Engine
}

func Start() {

	userMongoRepository := infra.NewUserMongoRepository()
	userConsumer := infra.NewConsumer("test_topic")

	userHandler := UserHandler{
		userService: domain.NewUserService(userMongoRepository),
	}

	authHandler := AuthHandler{
		authService: domain.NewAuthService(userMongoRepository),
	}

	r := Routes{
		router: gin.Default(),
	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.router.Group("/api")
	v1 := api.Group("/v1")

	users := v1.Group("/users")

	users.Use(authHandler.authMiddleware)
	//users.GET("/", userHandler.getAllUsers)
	users.GET("/:userId", userHandler.getUserByUserId)
	users.DELETE("/:userId", userHandler.deleteUser)
	users.PUT("/:userId", userHandler.updateUser)

	login := v1.Group("/login")
	login.POST("/", authHandler.handleLogin)

	signup := v1.Group("/signup")
	signup.POST("/", userHandler.createUser)

	auth := v1.Group("/auth")
	auth.GET("/", authHandler.isTokenValid)

	userConsumer.Start()
	go updatePurchaseHistory(*userConsumer, userMongoRepository)

	err := r.router.Run(":8089")
	if err != nil {
		logger.Fatal("Unable to start user service")
	}
}

func updatePurchaseHistory(userConsumer infra.Consumer, umr domain.UserMongoRepository) {
	for {
		oi := <-userConsumer.Output
		umr.UpdatePurchaseHistory(oi.UserId, oi.OrderId, oi.OrderAmount)
		logger.Info("Updated purchase history")
	}
}
