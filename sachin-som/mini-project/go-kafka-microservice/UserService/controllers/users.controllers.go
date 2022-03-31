package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kafka-microservice/UserService/models"
	"github.com/go-kafka-microservice/UserService/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var user models.User
	if err := gctx.ShouldBindJSON(&user); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := uc.UserService.CreateUser(&user); err != nil {
		gctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusCreated, gin.H{"message": "User Created."})
}

func (uc *UserControllers) GetUser(gctx *gin.Context) {
	var userId primitive.ObjectID
	var err error
	if userId, err = primitive.ObjectIDFromHex(gctx.Param("id")); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	user, err := uc.UserService.GetUser(userId)
	if err != nil {
		gctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"message": user})
}

func (uc *UserControllers) Login(gctx *gin.Context) {
	var credentials models.Credentials
	if err := gctx.ShouldBindJSON(&credentials); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if credentials.Email == "" || credentials.Password == "" {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": "Please Provide email and password."})
		return
	}
	token, err := uc.UserService.Login(&credentials)
	if err != nil {
		gctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	fmt.Println("*****************************")
	fmt.Println(token)
	fmt.Println("IN CONTROLLLERS")
	fmt.Println("**************************")
	gctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserControllers) RegisterUserRoutes(rg *gin.RouterGroup) {
	userGroup := rg.Group("/")
	userGroup.POST("/create", uc.CreateUser)
	userGroup.GET("/get/:id", uc.GetUser)
	userGroup.POST("/login", uc.Login)
}
