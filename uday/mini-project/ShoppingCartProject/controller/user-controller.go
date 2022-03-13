package controller

import (
	"fmt"

	"github.com/Udaysonu/SwiggyGoLangProject/entity"
	"github.com/Udaysonu/SwiggyGoLangProject/service"
	"github.com/gin-gonic/gin"
)

type UserController interface{
	SignUpUser(ctx *gin.Context)
	IsUserPresent(ctx *gin.Context)bool
	GetUser(ctx *gin.Context)entity.User
}

type userController struct{
	service service.UserService
}

func UserNewController(service service.UserService) UserController{
	return	&userController{
		service,
	}
}
type TempUser struct{
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    int `json:"phone"`
}
func (c *userController)SignUpUser(ctx *gin.Context){
	var user TempUser
	error:=ctx.ShouldBindJSON(&user)
	fmt.Println(user,error)
	c.service.SignUpUser(user.Username,user.Password,user.Email,user.Phone);
}

func (c *userController)IsUserPresent(ctx *gin.Context)bool{
	var user TempUser
	ctx.BindJSON(&user)
	fmt.Println(user)
	return c.service.IsUserPresent(user.Username,user.Password);
}

func (c *userController)GetUser(ctx *gin.Context)entity.User{
	var user TempUser
	ctx.BindJSON(&user)
	fmt.Println(user)
	return c.service.GetUser(user.Username,user.Password);
}
