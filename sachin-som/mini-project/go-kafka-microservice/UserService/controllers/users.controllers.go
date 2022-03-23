package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/UserService/services"
)

type UserControllers struct {
	UserService services.UserService
}

func NewUserControllers(userService services.UserService) *UserControllers {
	return &UserControllers{
		UserService: userService,
	}
}

func (uc *UserControllers) CreateUser(gctx *gin.Context) {
	gctx.JSON(200, nil)
}

func (uc *UserControllers) GetUser(gctx *gin.Context) {
	gctx.JSON(200, nil)
}

func (uc *UserControllers) RegisterUserRoutes(rg *gin.RouterGroup) {
	userGroup := rg.Group("/users")
	userGroup.POST("/create", uc.CreateUser)
	userGroup.GET("/get/:id", uc.GetUser)
}
