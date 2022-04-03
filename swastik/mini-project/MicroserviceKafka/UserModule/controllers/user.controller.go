package controllers

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swastiksahoo153/train-reservation-system/models"
	"github.com/swastiksahoo153/train-reservation-system/services"
	"github.com/swastiksahoo153/train-reservation-system/middlewares"
)

type UserController struct{
	UserService services.UserService
}

func New(userservice services.UserService) UserController{
	return UserController{
		UserService: userservice,
	}
}


// @Summary Register User
// @Description To register a new user for the app.
// @Tags User
// @Schemes
// @Accept json
// @Produce json
// @Param        user	body	models.User  true  "User structire"
// @Success	201  {string} 	success
// @Failure	400  {number} 	http.http.StatusBadRequest
// @Failure	502  {number} 	http.StatusBadGateway
// @Router /userRegistration [POST]
func (uc *UserController) RegisterUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return 
	}

	hashPassword := services.HashPassword(user.Password)
	user.Password = hashPassword

	err := uc.UserService.CreateUser(&user)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusCreated,  gin.H{"message": "success"})
}


// @Summary Login User
// @Description To login a new user for the app.
// @Tags User
// @Schemes
// @Accept json
// @Produce json
// @Param	user	body	models.Login  true  "User structire"
// @Success	200  {string} 	token
// @Failure	400  {number} 	http.http.StatusBadRequest
// @Failure	502  {number} 	http.StatusBadGateway
// @Router /userLogin [POST]
func (uc *UserController) LoginUser(ctx *gin.Context)  {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return 
	}
	fmt.Println("username", user.Username)
	foundUser, err := uc.UserService.GetUser(&user.Username)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	if foundUser.Username == ""{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"user not found"})
	}

	passwordIsValid, msg := services.VerifyPassword(user.Password, foundUser.Password)
	if !passwordIsValid{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	token, err := services.CreateToken(foundUser.Username, foundUser.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": token})
}


// @Summary Get User
// @Description To get user details.
// @Tags User
// @Schemes
// @Accept json
// @Param username path string true "User Name"
// @Produce json
// @Success	200  {object} 	models.User
// @Failure	502  {number} 	http.StatusBadGateway
// @Security Bearer Token
// @Router /user/get/{username} [GET]
func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("username")
	fmt.Println("username", username)
	user, err := uc.UserService.GetUser(&username)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK,  user)
}


// @Summary Get all User details
// @Description To get every user detail.
// @Tags User
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {array} 	models.User
// @Failure	502  {number} 	http.StatusBadGateway
// @Security Bearer Token
// @Router /user/getall [GET]
func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.UserService.GetAll()
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK, users)
}


// @Summary Update User
// @Description To get user details.
// @Tags User
// @Schemes
// @Accept json
// @Param username path string true "User Name"
// @Produce json
// @Success	200  {string} 	success
// @Failure	502  {number} 	http.StatusBadGateway
// @Security Bearer Token
// @Router /user/update/{username} [PATCH]
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return 
	}
	err := uc.UserService.UpdateUser(&user)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK,  gin.H{"message": "success"})
}


// @Summary Delete User
// @Description To remove a particular user.
// @Tags User
// @Schemes
// @Accept json
// @Param username path string true "User Name"
// @Produce json
// @Success	200  {string} 	success
// @Failure	502  {number} 	http.StatusBadGateway
// @Security Bearer Token
// @Router /user/delete/{username} [DELETE]
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("username")
	err := uc.UserService.DeleteUser(&username)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK,  gin.H{"message": "success"})
}


// Routes
func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	public := rg.Group("")
	public.POST("/userRegistration",uc.RegisterUser)
	public.POST("/userLogin",uc.LoginUser)
	userroute := rg.Group("")
	userroute.Use(middlewares.AuthenticateJWT())
	userroute.GET("/user/get/:username", uc.GetUser)
	userroute.GET("/user/getall", uc.GetAll)
	userroute.PATCH ("/user/update", uc.UpdateUser)
	userroute.DELETE("/user/delete/:username", uc.DeleteUser)
}
