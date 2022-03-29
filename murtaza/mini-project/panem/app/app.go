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

	// firstName := "Murtaza"
	// lastName := "Sadriwala"
	// phone := "9900887766"
	// email := "murtaza896@gmail.com"
	// username := "murtaza896"
	// password := "Pass!23"
	// role := domain.Admin

	// user, _ := userHandler.userService.CreateUser(firstName, lastName, phone, email, username, password, role)
	// user, _ := userHandler.userService.CreateUserInMongo(firstName, lastName, phone, email, username, password, role)
	// userPersistedEntity, _ := userRepository.FindByEmail(user.Email())
	// fmt.Println(userPersistedEntity)
	// fmt.Println(user)

	r := Routes{
		router: gin.Default(),
	}

	v1 := r.router.Group("/v1")

	users := v1.Group("/users")

	users.GET("/", userHandler.demoHandlerFunc)
	users.GET("/:userId", userHandler.getUserByUserId)
	users.POST("/", userHandler.createUser)
	users.DELETE("/:userId", userHandler.deleteUser)
	// users.PUT("/:userId", demoHandlerFunc)
	// users.POST("/signup", demoHandlerFunc)
	// users.POST("/login", demoHandlerFunc)

	r.router.Run(":8089")
}
