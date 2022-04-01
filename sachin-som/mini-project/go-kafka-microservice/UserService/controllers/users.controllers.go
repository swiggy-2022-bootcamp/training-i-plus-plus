package controllers

import (
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

// CreateUser godoc
// @Summary      Create New User
// @Description  user creation API
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user body requests.UserRequest  true "user request structure"
// @Success      200  {object} 	responses.MessageResponse
// @Failure      400  {object} 	responses.MessageResponse
// @Failure      500  {object} 	responses.MessageResponse
// @Failure      502  {object} 	responses.MessageResponse
// @Router       /create [post]
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

// GetUser godoc
// @Summary      Get User By Id
// @Description  user get API
// @Tags         User
// @Accept       json
// @Produce      json
// @securityDefinitions.apikey ApiKeyAuth
// @Param        id path string  true "user id"
// @Success      200  {object} 	responses.UserResponse
// @Failure      400  {object} 	responses.MessageResponse
// @Failure      500  {object} 	responses.MessageResponse
// @Failure      502  {object} 	responses.MessageResponse
// @Router       /get/{id} [get]
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

// Login godoc
// @Summary      Login User
// @Description  login user API
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        credentials body requests.LoginRequest  true "user credentials structure"
// @Success      200  {object} 	responses.TokenResponse
// @Failure      400  {object} 	responses.MessageResponse
// @Failure      500  {object} 	responses.MessageResponse
// @Failure      502  {object} 	responses.MessageResponse
// @Router       /login [post]
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
	gctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserControllers) RegisterUserRoutes(rg *gin.RouterGroup) {
	userGroup := rg.Group("/")
	userGroup.POST("/create", uc.CreateUser)
	userGroup.GET("/get/:id", uc.GetUser)
	userGroup.POST("/login", uc.Login)
}
