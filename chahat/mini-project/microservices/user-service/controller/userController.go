package controller
import (
	"context"
	"fmt"
	"log"
	// "strconv"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	helper "bhatiachahat/user-service/helper"
	model "bhatiachahat/user-service/model"
//	productmodel "bhatiachahat/product-service/model"
	database "bhatiachahat/user-service/db"
//	productusermodels "bhatiachahat/user-service/model"
"bhatiachahat/user-service/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
var validate = validator.New()

func HashPassword(password string) string{
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err!=nil{
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string)(bool, string){
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err!= nil {
		msg = fmt.Sprintf("email of password is incorrect")
		check=false
	}
	return check, msg
}

// SignUpUser godoc
// @Summary To register a new user in the online shopping application
// @Description This request will register a new user.
// @Tags User
// @Schemes
// @Accept json
// @Produce json
// @Param req body model.User true "User details"
// @Success	201  {object} 	model.User
// @Failure	400  {number} 	http.http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Router /user/signup [POST]
func Signup()gin.HandlerFunc{

	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user model.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// validationErr := validate.Struct(user)
		// if validationErr != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error":validationErr.Error()})
		// 	return
		// }

		count, err := userCollection.CountDocuments(ctx, bson.M{"email":user.Email})
		defer cancel()
		if err != nil {
			
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the email"})
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		count, err = userCollection.CountDocuments(ctx, bson.M{"phone":user.Phone})
		defer cancel()
		if err!= nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the phone number"})
		}

		if count >0{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"this email or phone number already exists"})
			return
		}else{
			user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			user.ID = primitive.NewObjectID()
			 user.User_id = user.ID.Hex()
			token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.Firstname, *user.Lastname, *user.Usertype,user.User_id)
			user.Token = &token
			user.Refresh_token = &refreshToken
			//var  usercart  []model.ProductUser
			usercart := []model.ProductUser{}
		//	var usercart = model.ProductUser
			user.UserCart= usercart
	
			result, insertErr := userCollection.InsertOne(ctx, user)
			if insertErr !=nil {
				
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			defer cancel()
			//c.JSON(http.StatusOK, resultInsertionNumber)
			c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
		}

	
	}

}

// LoginUser godoc
// @Summary User gets logged in.
// @Description This request will login a user.
// @Tags User
// @Schemes
// @Accept json
// @Produce json
// @Param        Password 	body	string  true  "password"
// @Param        Email 		body	string  true  "email id"
// @Success	200  {string} 	token
// @Failure	400  {number} 	http.http.StatusBadRequest
// @Failure	404  {number} 	http.http.StatusNotFound
// @Failure	500  {number} 	http.StatusInternalServerError
// @Router /users/login [POST]
func Login() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user model.User
		var foundUser model.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return 
		}
		count, err := userCollection.CountDocuments(ctx,  bson.M{"email":user.Email})
		defer cancel()
		if err!= nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while login "})
		}

		if count <1{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Please register first"})
			return
		}

		nerr := userCollection.FindOne(ctx, bson.M{"email":user.Email}).Decode(&foundUser)
		defer cancel()
		if nerr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"email or password is incorrect"})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if passwordIsValid != true{
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		if foundUser.Email == nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"user not found"})
		}
		token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.Firstname, *foundUser.Lastname, *foundUser.Usertype, foundUser.User_id)
		helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)
		err = userCollection.FindOne(ctx, bson.M{"user_id":foundUser.User_id}).Decode(&foundUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, foundUser)
	}
}

func AddToCart() gin.HandlerFunc{
	return func(c *gin.Context) {
		userId := c.Param("user_id")
		//fmt.Println(userId)

				var ctx, cancel = context.WithTimeout(context.Background(), 500*time.Second)
		
				//var user model.User
			

				// user,err := userCollection.FindOne(ctx, bson.M{"user_id":userId})
				// defer cancel()
				// if err != nil{
				// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				// 	return
				// }
				// fmt.Println(user.UserCart)
				var product model.ProductUser
				if err := c.BindJSON(&product); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
					return 
				}
				var user model.User
			///	err = userCollection.FindOne(ctx, bson.M{"user_id":foundUser.User_id}).Decode(&foundUser)
				 userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&user)
	fmt.Println(user,"\n")
	fmt.Println("----------------");
			//	fmt.Println(product)
			//	filterInventory := bson.D{bson.E{Key: "user_id", Value: userId}}
			//	fmt.Println(filterInventory)
	pushProuductID := bson.D{bson.E{Key: "$push", Value: bson.D{bson.E{Key: "user_cart", Value: product}}}}
	defer cancel()
	fmt.Print("push",pushProuductID)

	// Push ProuductID to inventory
	
	//result,err := userCollection.UpdateOne(ctx,user, pushProuductID)
	// if err!=nil{
	// 	fmt.Println(err)
	// }
	// fmt.Println("res",result)
	result, err := userCollection.UpdateOne(ctx,bson.M{"user_id":userId},pushProuductID)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//get updated user details
	var updatedproduct model.User
	if result.MatchedCount == 1 {
		err := userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&updatedproduct)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, updatedproduct)    
//}}

	

	// filterInventory := bson.D{bson.E{Key: "user_id", Value: userId}}
	// pushProuduct := bson.D{bson.E{Key: "$push", Value: bson.D{bson.E{Key: "user_cart", Value: product}}}}

	// // Craete Product and get prouduct ID
	
	// // Push ProuductID to inventory
	// res, _:= userCollection.UpdateOne(ctx, filterInventory, pushProuduct)
	// fmt.Println((res))

	}
}
func GetCart() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id")
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user model.User
		err := userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&user)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		usercart := []model.ProductUser{}
					usercart=user.UserCart
		c.JSON(http.StatusOK,usercart)

	}
}

// GetAllUsers godoc
// @Summary Get all Users list.
// @Description Get details of all Users.
// @Tags GeneralUser
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {array} 	model.User
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /users [GET]
func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []model.User
		defer cancel()

		results, err :=userCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser model.User
			if err = results.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

		users = append(users, singleUser)
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}},
		)
	}
}

// GetUser godoc
// @Summary Get User by ID.
// @Description View individual user details.
// @Tags User
// @Schemes
// @Param id path string true "User id"
// @Accept json
// @Produce json
// @Success	200  {object} 	model.User
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /users/{id} [GET]
func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id")

		// if err := helper.MatchUserTypeToUid(c, userId); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		// 	return
		// }
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user model.User
		err := userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&user)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data":user}})
	}
}


// DeleteUserByID godoc
// @Summary Delete User by ID.
// @Description User can delete his account.
// @Tags User
// @Schemes
// @Param id path string true "User id"
// @Accept json
// @Produce json
// @Success	200  {string} 	User successfully deleted!
// @Failure	404  {number} 	http.http.StatusNotFound
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /users/{id} [DELETE]
func DeleteUser()gin.HandlerFunc{

	return func(c *gin.Context){
		userId := c.Param("user_id")
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		
		res,err := userCollection.DeleteOne(ctx, bson.M{"user_id":userId})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		if res.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}},
			)
			return
		}
		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}},
		)

	}

}

// EditUserByID godoc
// @Summary Edit User by ID.
// @Description Edit details of a User.
// @Tags GeneralUser
// @Schemes
// @Param id path string true "User id"
// @Accept json
// @Produce json
// @Param req body model.User true  "User details"
// @Success	200  {object} 	model.User
// @Failure	400  {number} 	http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /users/{id} [PUT]
func EditUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        userId := c.Param("user_id")
        var user model.User
        defer cancel()
	    // validate the request body
        if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data":err.Error()}})
            return
        }

        update := bson.M{"first_name": user.Firstname, "last_name": user.Lastname, "phone":user.Phone}
		
	
        result, err := userCollection.UpdateOne(ctx,bson.M{"user_id":userId}, bson.M{"$set": update})
	
        if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
        }

        //get updated user details
        var updateduser model.User
        if result.MatchedCount == 1 {
            err := userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&updateduser)
            if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
            }
        }

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updateduser}})   }}


