package auth

import (
	"context"
	"fmt"
	"net/http"
	JWTManager "swiggy/gin/lib/helpers"
	db "swiggy/gin/lib/utils"
	"swiggy/gin/services/user"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func init() {
	db.ConnectDB()
	JWTManager.NewJWTManager("Ishan", time.Hour*50)
}

func Login(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	body := LoginBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user := &user.User{}
	if err := db.DataStore.Collection("user").FindOne(ctx, bson.M{"username": body.Username}).Decode(&user); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User Not found"})
		return
	}

	if user == nil || !user.IsCorrectPassword(body.Password) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Password Incorrect"})
		return
	}

	token, err := JWTManager.Manager.Generate(user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error while creating token"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"user": user, "token": token})
}

func Signup(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	body := SignUpBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := user.NewUser(body.Username, body.Password, body.Role)

	res, err := db.DataStore.Collection("user").InsertOne(ctx, user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("can not convert to oid %v", err)})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"username": body.Username, "Id": oid.Hex()})
}

func CheckAuth(c *gin.Context) {
	user := c.GetString("User")

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Auth Working", "user": user})
}
