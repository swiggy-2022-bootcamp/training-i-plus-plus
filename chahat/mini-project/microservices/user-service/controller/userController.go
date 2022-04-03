package productcontroller
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
	
			resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
			if insertErr !=nil {
				msg := fmt.Sprintf("User item was not created")
				c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
				return
			}
			defer cancel()
			c.JSON(http.StatusOK, resultInsertionNumber)
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

// // GetAllUsers godoc
// // @Summary Get all users.
// // @Description Get all users.
// // @Tags User
// // @Schemes
// // @Accept json
// // @Produce json
// // @Success	200  {array} 	model.User
// // @Failure	500  {number} 	http.StatusInternalServerError
// // @Security Bearer Token
// // @Router /users [GET]
// func GetUsers() gin.HandlerFunc{
// 	return func(c *gin.Context){
// 		if err := helper.CheckUserType(c, "ADMIN"); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
// 			return
// 		}
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
// 		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
// 		if err != nil || recordPerPage <1{
// 			recordPerPage = 10
// 		}
// 		page, err1 := strconv.Atoi(c.Query("page"))
// 		if err1 !=nil || page<1{
// 			page = 1
// 		}

// 		startIndex := (page - 1) * recordPerPage
// 		startIndex, err = strconv.Atoi(c.Query("startIndex"))

// 		matchStage := bson.D{{"$match", bson.D{{}}}}
// 		groupStage := bson.D{{"$group", bson.D{
// 			{"_id", bson.D{{"_id", "null"}}}, 
// 			{"total_count", bson.D{{"$sum", 1}}}, 
// 			{"data", bson.D{{"$push", "$$ROOT"}}}}}}
// 		projectStage := bson.D{
// 			{"$project", bson.D{
// 				{"_id", 0},
// 				{"total_count", 1},
// 				{"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},}}}
// result,err := userCollection.Aggregate(ctx, mongo.Pipeline{
// 	matchStage, groupStage, projectStage})
// defer cancel()
// if err!=nil{
// 	c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while listing user items"})
// }
// var allusers []bson.M
// if err = result.All(ctx, &allusers); err!=nil{
// 	log.Fatal(err)
// }
// c.JSON(http.StatusOK, allusers[0])}}

// // GetUserByID godoc
// // @Summary Get User by ID.
// // @Description View individual user details.
// // @Tags User
// // @Schemes
// // @Param id path string true "User id"
// // @Accept json
// // @Produce json
// // @Success	200  {object} 	model.User
// // @Failure	500  {number} 	http.StatusInternalServerError
// // @Security Bearer Token
// // @Router /users/{id} [GET]
// func GetUser() gin.HandlerFunc{
// 	return func(c *gin.Context){
// 		userId := c.Param("user_id")

// 		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
// 			return
// 		}
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 		var user model.User
// 		err := userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&user)
// 		defer cancel()
// 		if err != nil{
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, user)
// 	}
// }


// // DeleteUserByID godoc
// // @Summary Delete User by ID.
// // @Description User can delete his account.
// // @Tags User
// // @Schemes
// // @Param id path string true "User id"
// // @Accept json
// // @Produce json
// // @Success	200  {string} 	User successfully deleted!
// // @Failure	404  {number} 	http.http.StatusNotFound
// // @Failure	500  {number} 	http.StatusInternalServerError
// // @Security Bearer Token
// // @Router /users/{id} [DELETE]
// func DeleteProduct()gin.HandlerFunc{

// 	return func(c *gin.Context){
// 		productId := c.Param("product_id")
		
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		
// 		res,err := productCollection.DeleteOne(ctx, bson.M{"product_id":productId})
// 		defer cancel()
// 		if err != nil{
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK,res)

// 	}

// }

// func DeleteGeneralUserByID() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		id := c.Param("id")
// 		defer cancel()

// 		objId, _ := primitive.ObjectIDFromHex(id)

// 		result, err := generalUserCollection.DeleteOne(ctx, bson.M{"_id": objId})

// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
// 			return
// 		}

// 		if result.DeletedCount < 1 {
// 			c.JSON(http.StatusNotFound,
// 				responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}},
// 			)
// 			return
// 		}

// 		c.JSON(http.StatusOK,
// 			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}},
// 		)
// 	}
// }


