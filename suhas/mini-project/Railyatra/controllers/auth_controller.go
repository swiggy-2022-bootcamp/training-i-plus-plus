package controllers

import (
	"context"
	"fmt"
	"gin-mongo-api/models"
	"gin-mongo-api/repository"
	"gin-mongo-api/responses"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//var registerCollection *mongo.Collection = config.GetCollection(config.DB, "register")

var registerrepo repository.AuthRepository

var (
	mySigningKey = []byte("secret")
)

func GetJWT(group string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	fmt.Print("p1")
	claims := token.Claims.(jwt.MapClaims)
	fmt.Print("p2")
	claims["authorized"] = true
	claims["group"] = group //group should be USER or ADMIN
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	fmt.Print("p3")
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

// GeneratePasswordHash handles generating password hash
// bcrypt hashes password of type byte
func GeneratePasswordHash(password []byte) string {
	// default cost is 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	// If there was an error panic
	if err != nil {
		panic(err)
	}

	// return stringified password
	return string(hashedPassword)
}

// PasswordCompare handles password hash compare
func PasswordCompare(password []byte, hashedPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)

	return err
}

func CheckUserEmail(email string) bool {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := userrepo.ReadAll()

	if err != nil {
		return false
	}

	//reading from the db in an optimal way
	// defer results.Close(ctx)
	// for results.Next(ctx) {
	// 	var singleUser models.User
	// 	if err = results.Decode(&singleUser); err != nil {
	// 		return false
	// 	}
	// 	if singleUser.Email == email {
	// 		return true
	// 	}
	// }
	for _, res := range results {
		if res.Email == email {
			return true
		}
	}
	return false
}

func CheckAdminEmail(email string) bool {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := adminrepo.ReadAll()

	if err != nil {
		return false
	}

	for _, res := range results {
		if res.Email == email {
			return true
		}
	}
	return false
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var register models.Register
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&register); err != nil {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&register); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newRegister := models.Register{
			Username: register.Username,
			Email:    register.Email,
			Group:    register.Group,
			Password: register.Password,
		}

		if register.Group == "ADMIN" || register.Group == "USER" {
			//result, err := registerCollection.InsertOne(ctx, newRegister)
			result, err := registerrepo.Insert(newRegister)
			if err != nil {
				c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			c.JSON(http.StatusCreated, responses.LoginResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"Inserted data": result}})
			return
		} else {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error not valid group"})
			return
		}

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var register models.Register
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&register); err != nil {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error in binding", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := avalidate.Struct(&register); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusBadRequest, Message: "error in validate register", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		if register.Group == "ADMIN" {
			//var admin_reg models.Register
			//err := registerCollection.FindOne(ctx, bson.M{"username": register.Username}).Decode(&admin_reg)
			admin_reg, err := registerrepo.Read(register.Username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.LoginResponse{Status: http.StatusInternalServerError, Message: "error in locating user", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			err = PasswordCompare([]byte(admin_reg.Password), []byte(register.Password))
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.LoginResponse{Status: http.StatusBadRequest, Message: "passowrd not correct", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			token, err := GetJWT("ADMIN")
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.LoginResponse{Status: http.StatusInternalServerError, Message: "error in generating token", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusCreated, responses.LoginResponse{Status: http.StatusCreated, Message: "success", Token: token})
			return
		} else if register.Group == "USER" {
			//var user_reg models.Register
			//err := registerCollection.FindOne(ctx, bson.M{"username": register.Username}).Decode(&user_reg)
			user_reg, err := registerrepo.Read(register.Username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.LoginResponse{Status: http.StatusInternalServerError, Message: "error in locating user", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			err = PasswordCompare([]byte(user_reg.Password), []byte(register.Password))
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.LoginResponse{Status: http.StatusBadRequest, Message: "passowrd not correct", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			token, err := GetJWT("USER")
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.LoginResponse{Status: http.StatusInternalServerError, Message: "error in generating token", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			c.JSON(http.StatusCreated, responses.LoginResponse{Status: http.StatusCreated, Message: "success", Token: token})
			return
		} else {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{Status: http.StatusInternalServerError, Message: "error not valid group"})
			return
		}

	}
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func IsAuthorized(group string) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.Request.Header.Get("Authorization")
		//normally Authorization the_token_xxx
		strArr := strings.Split(bearToken, " ")
		if len(strArr) != 2 {
			respondWithError(c, 401, "No bearer token")
			return
		}
		if !VerifyClaims(strArr[1], group) {
			respondWithError(c, 401, "Functionality not available for this user")
			return
		}
		token, err := jwt.Parse(strArr[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(("invalid Signing Method"))
			}

			return mySigningKey, nil
		})
		if err != nil {
			respondWithError(c, 501, err.Error())
			return
		}
		if !token.Valid {
			respondWithError(c, 401, "Invalid token")
			return
		}
		c.Next()
	}
}

func VerifyClaims(tokenStr string, group string) bool {

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(("invalid Signing Method"))
		}
		return mySigningKey, nil
	})

	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["group"] == group {
			return true
		} else {
			return false
		}
	} else {
		log.Printf("Invalid JWT Token")
		return false
	}
}
