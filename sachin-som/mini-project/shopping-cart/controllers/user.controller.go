package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sachinsom93/shopping-cart/services"
)

type UserController struct {
	UserService services.UserServices
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) GetAllUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/users")
	userRoute.POST("/create", uc.CreateUser)
	userRoute.GET("/get/:email", uc.GetUser)
	userRoute.GET("/getall", uc.GetAllUser)
	userRoute.PATCH("/update", uc.UpdateUser)
	userRoute.DELETE("/delete:email", uc.DeleteUser)
}
