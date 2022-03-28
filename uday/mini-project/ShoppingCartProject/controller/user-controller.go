package controller

import (
	"fmt"

	"github.com/Udaysonu/SwiggyGoLangProject/entity"
	"github.com/Udaysonu/SwiggyGoLangProject/service"
	"github.com/gin-gonic/gin"
)

// type UserController interface{
// 	SignUpUser(ctx *gin.Context)
// 	IsUserPresent(ctx *gin.Context)bool
// 	GetUser(ctx *gin.Context)entity.User
// }

// type userController struct{
// 	service service.UserService
// }

// func UserNewController(service service.UserService) UserController{
// 	return	&userController{
// 		service,
// 	}
// }
type TempUser struct{
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    int `json:"phone"`
	Location int `json:"location"`
}

// SignUpUser godoc
// @Summary Sign Up User
// @Description get string by ID
// @Consume application/x-www-form-urlencoded
// @Accept json
// @Produce json
// @Param newuser body TempUser true "New User Details"
// @Success 200 {object} entity.User
// @Failure 400 {object} entity.User
// @Failure 500 {object} entity.User
// @Router /user/signuser [post]
func SignUpUser(ctx *gin.Context)entity.User{
	var user TempUser
	error:=ctx.ShouldBindJSON(&user)
	fmt.Println(user,error)
	return service.SignUpUser(user.Username,user.Password,user.Email,user.Phone,user.Location);
}

func  IsUserPresent(ctx *gin.Context)bool{
	var user TempUser
	ctx.BindJSON(&user)
	fmt.Println(user)
	return service.IsUserPresent(user.Username,user.Password);
}

func  GetUser(ctx *gin.Context)entity.User{
	var user TempUser
	ctx.BindJSON(&user)
	fmt.Println(user)
	return service.GetUser(user.Username,user.Password);
}
