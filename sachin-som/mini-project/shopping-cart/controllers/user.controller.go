package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sachinsom93/shopping-cart/models"
	"github.com/sachinsom93/shopping-cart/services"
)

type UserController struct {
	UserService services.UserServices
}

func NewUserController(userService services.UserServices) UserController {
	return UserController{
		UserService: userService,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "success"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	email := ctx.Param("email")
	user, err := uc.UserService.GetUser(&email)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAllUser(ctx *gin.Context) {
	users, err := uc.UserService.GetAllUser()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	email := ctx.Param("email")
	err := uc.UserService.DeleteUser(&email)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/users")
	userRoute.POST("/create", uc.CreateUser)
	userRoute.GET("/get/:email", uc.GetUser)
	userRoute.GET("/getall", uc.GetAllUser)
	userRoute.PATCH("/update", uc.UpdateUser)
	userRoute.DELETE("/delete/:email", uc.DeleteUser)
}
