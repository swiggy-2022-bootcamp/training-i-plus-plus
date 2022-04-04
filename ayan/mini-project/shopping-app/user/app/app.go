package app

import (
	"fmt"
	"user/db"
	"user/docs"
	"user/domain"
	"user/utils/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start() {

	dbClient := db.NewDbClient()

	// userRepo := db.NewUserRepositoryListDB([]db.User{})
	userRepo := db.NewUserRepositoryDB(dbClient)
	userService := domain.NewUserService(userRepo)
	userHandlers := UserHandlers{service: userService}

	userRouter := gin.Default()

	apiRouter := userRouter.Group("/api")

	docs.SwaggerInfo.BasePath = "/api"
	apiRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRoutesGroup := apiRouter.Group("/users")

	userRoutesGroup.GET("/", userHandlers.HelloWorldHandler)
	userRoutesGroup.GET("/:userEmail", userHandlers.GetUserByEmail)
	userRoutesGroup.POST("/register", userHandlers.Register)
	userRoutesGroup.POST("/login", userHandlers.Login)
	userRoutesGroup.PUT("/update", userHandlers.UpdateUser)
	userRoutesGroup.DELETE("/:userEmail", userHandlers.DeleteUserByEmail)
	userRoutesGroup.POST("/verifyToken", userHandlers.VerifyUserToken)

	userRouter.Run(":8080")
	logger.Info(fmt.Sprintf("Starting server on %s:%s ...", "127.0.0.1", "8080"))
}
