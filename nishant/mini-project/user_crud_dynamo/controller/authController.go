package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/user_crud_dynamo/config"
	"github.com/swiggy-2022-bootcamp/training-i-plus-plus/nishant/mini-project/user_crud_dynamo/model"
	"golang.org/x/crypto/bcrypt"
)

var secret []byte = []byte(config.C().JWT.Secret)

type UserCreds struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserAuthResponse struct {
	UserId string `json:"userId"`
	Token  string `json:"token"`
}

type Claims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

func (cont Controller) CreateToken(c *gin.Context) {
	userId := c.GetString("userId")

	claims := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(5 * time.Hour).Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(secret)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		log.Println("Error while generation jwt " + err.Error())
		c.AbortWithStatusJSON(500, gin.H{
			"error": "error while creating token",
		})
		return
	}

	c.JSON(200, UserAuthResponse{
		userId,
		tokenString,
	})
}

func (cont Controller) VerifyToken(c *gin.Context) {

	if len(c.Request.Header["Authorization"]) == 0 {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "Authorization token not found",
		})
		return
	}

	tknStr := c.Request.Header["Authorization"][0]
	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("userId", claims.UserId)
	log.Println("User " + claims.UserId + " Autenticated!!")
	c.Next()
}

// Login godoc
// @Summary User Log In
// @Description get auth token from username password
// @Tags Users
// @Param   user      body UserCreds true  "user creds"
// @Accept  json
// @Success 200		{object} UserAuthResponse
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /login [post]
func (cont Controller) Login(c *gin.Context) {
	cred := UserCreds{}

	if err := c.BindJSON(&cred); err != nil {
		c.Error(err)
		return
	}
	log.Printf("cred %+v", cred)

	filt := expression.Name("email").Equal(expression.Value(cred.Email))
	proj := expression.NamesList(expression.Name("userId"), expression.Name("pass"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		log.Println("Expression Error : " + err.Error())
	}

	result, err := cont.Dynamo.Scan(&dynamodb.ScanInput{
		TableName:                 aws.String(model.UserTableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		Limit:                     aws.Int64(1),
	})

	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	var resList []struct {
		UserId string `json:"userId,omitempty"`
		Pass   []byte `json:"pass,omitempty""`
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &resList)
	if err != nil {
		log.Panicln(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	log.Printf("fetch result %+v", result)

	if len(resList) == 0 {
		c.AbortWithStatusJSON(404, gin.H{
			"msg": "user not found",
		})
		return
	}

	if err = bcrypt.CompareHashAndPassword(resList[0].Pass, []byte(cred.Password)); err != nil {
		//invalid password
		log.Println("Invalid password")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("userId", resList[0].UserId)
	log.Println("User " + resList[0].UserId + " Autenticated!!")
	c.Next()
}
