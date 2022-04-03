package controllers

import (
	log "authService/logger"
	"authService/models"
	"authService/repository"
	"authService/responses"
	"context"
	"net/http"
	"strings"
	"time"

	pb "authService/protobuf"

	"fmt"
	"net"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

//var registerCollection *mongo.Collection = config.GetCollection(config.DB, "register")

var registerrepo repository.AuthRepository

var (
	mySigningKey = []byte("secret")
	avalidate    = validator.New()
)
var (
	errLog = log.ErrorLogger.Println
)

type Server struct {
	pb.UnimplementedAuthenticationServiceServer
}

func GetJWT(group string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
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
func PasswordCompare(hashedPassword []byte, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)

	return err
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

		hash := GeneratePasswordHash([]byte(register.Password))

		newRegister := models.Register{
			Username: register.Username,
			Email:    register.Email,
			Group:    register.Group,
			Password: hash,
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
			fmt.Println(admin_reg.Password, register.Password)
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

// Rest api
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
			fmt.Println("Invalid signing method")
			return nil, fmt.Errorf(("invalid Signing Method"))
		}
		return mySigningKey, nil
	})

	if err != nil {
		fmt.Println(err)
		fmt.Println("No auth for this token")
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["group"])
		if claims["group"] == group {
			return true
		} else {
			return false
		}
	} else {
		errLog("Invalid JWT Token")
		return false
	}
}

// grpc

func (s *Server) Authenticate(ctx context.Context, in *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	bearerToken := in.Token
	fmt.Println(bearerToken)

	if !VerifyClaims(bearerToken, in.Group) {
		errLog("Inavlid Group for this token")
		return &pb.AuthenticateResponse{
			Confirmation: false,
			Message:      "Inavlid Group for this token",
		}, nil
	}

	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(("invalid Signing Method"))
		}

		return mySigningKey, nil
	})
	if err != nil {
		errLog(err.Error())
		return &pb.AuthenticateResponse{
			Confirmation: false,
			Message:      err.Error(),
		}, nil
	}
	if !token.Valid {
		errLog("Invalid token")
		return &pb.AuthenticateResponse{
			Confirmation: false,
			Message:      "Invalid token",
		}, nil
	}

	return &pb.AuthenticateResponse{
		Confirmation: true,
		Message:      "Success",
	}, nil
}

func Startgrpc() error {
	gr := grpc.NewServer()
	lis, err := net.Listen("tcp", ":6010")
	if err != nil {
		errLog("Failed to listen: %v", err)
		fmt.Printf("Failed to listen: %v\n", err)
		return err
	}

	pb.RegisterAuthenticationServiceServer(gr, &Server{})
	err = gr.Serve(lis)
	if err != nil {
		errLog("Failed to serve: %v", err)
		fmt.Printf("Failed to serve: %v\n", err)
		return err
	}
	return nil
}
