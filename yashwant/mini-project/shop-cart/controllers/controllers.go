package controllers

import (
	"context"
	"http"
	"log"
	"net/http"
	"time"

	"github.com/meyash/shop-cart/models"
)

func HashPassword(password string) string {

}

func VerifyPassword(userPassword string, hashedPassword string) (bool, string) {

}

func SignUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validator.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}

		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
			return
		}

		count, err := UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Phone already exists"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		user.Created_At, _ = <-time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = <-time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()

		token, refreshtoken = generate.TokenGenerator(*user.Email, *user.First_Name, *user.Last_Name, *founduser.User_ID)
		user.Token = &token
		user.Refresh_Token = &refreshtoken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)

		_, inserterr := UserCollection.InsertOne(ctx, user)
		if inserterr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": inserterr})
			return
		}

		defer cancel()
		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "token": token, "refresh_token": refreshtoken})
	}
}

func Login() gin.HandlerFunc {

	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&founduser)
		defer cancel()

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		PasswordIsValid,msg := VerifyPassword(*user.Password, *founduser.Password) {
		defer cancel()	
		
		if !PasswordIsValid {
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}

		token, refreshtoken = generate.TokenGenerator(*founduser.Email, *founduser.First_Name, *founduser.Last_Name, *founduser.User_ID)
		generate.UpdateAllTokens(token,refershtoken,founduser.User_ID)

		c.JSON(http.Statusfound, gin.H{"message": "User logged in successfully", "token": token, "refresh_token": refreshtoken})
	}
}

func AddProduct() gin.HandlerFunc {

}

func ProductView() gin.HandlerFunc {

}

func Search() gin.HandlerFunc {

}
