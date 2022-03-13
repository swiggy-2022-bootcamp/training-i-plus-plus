package middleware
import (
	 "github.com/dgrijalva/jwt-go"
	 "github.com/gin-gonic/gin" 
	 "github.com/Udaysonu/SwiggyGoLangProject/service"
	 "github.com/Udaysonu/SwiggyGoLangProject/entity"
	 "time"
	 "fmt"
)
var jwtKey=[]byte("sercret_key")

type Credentials struct{
	Username string `json:"username"`
	Password string	`json:"password"`
}

func Login(ctx *gin.Context,service service.UserService)string{
	var credentials Credentials
	ctx.BindJSON(&credentials)
	fmt.Println(credentials)
	var user entity.User=service.GetUser(credentials.Username,credentials.Password)
	fmt.Println(user)
	if user.Username==credentials.Username{
		expirationTime:=time.Now().Add(time.Minute*5)
		claims:=&Claims{
			Username: credentials.Username,
			Password: credentials.Password,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt:expirationTime.Unix(),
			},
		}
		tokem:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
		fmt.Println(tokem)
		tokenString,_:=tokem.SignedString(jwtKey)
		return tokenString
	} else{
		return "error in creating token! check credentials"
	}
}

func CheckAuth(ctx *gin.Context,service service.UserService)bool{
 	 
	tokenStr:=ctx.GetHeader("token")
	fmt.Println(tokenStr)
	claims:=&Claims{}
	tkn,err:=jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token)(interface{},error){
			return jwtKey,nil
		})
	if err!=nil{
		if err==jwt.ErrSignatureInvalid{
			return false
		}
	}
	if !tkn.Valid{
			return false
	}
	fmt.Println("welcome user")
	return true
}
type Claims struct{
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}