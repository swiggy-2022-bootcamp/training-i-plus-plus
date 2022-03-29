package app

import (
	"user/db"
	"user/domain"

	"github.com/gin-gonic/gin"
)

func Start() {

	dbClient := db.NewDbClient()

	// userRepo := db.NewUserRepositoryListDB([]db.User{})
	userRepo := db.NewUserRepositoryDB(dbClient)
	userService := domain.NewUserService(userRepo)
	userHandlers := UserHandlers{service: userService}

	userRouter := gin.Default()

	userRoutesGroup := userRouter.Group("/users")

	// swagger:operation GET /users
	userRoutesGroup.GET("/", userHandlers.HelloWorldHandler)
	userRoutesGroup.GET("/:userEmail", userHandlers.GetUserByEmail)
	userRoutesGroup.POST("/register", userHandlers.Register)
	userRoutesGroup.POST("/login", userHandlers.Login)
	userRoutesGroup.PUT("/update", userHandlers.UpdateUser)
	userRoutesGroup.DELETE("/:userEmail", userHandlers.DeleteUserByEmail)

	userRouter.Run(":8080")
}
