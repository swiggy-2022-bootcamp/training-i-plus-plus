package controller

import (
	"fmt"
	"strconv"
	"github.com/Udaysonu/SwiggyGoLangProject/entity"
	"github.com/Udaysonu/SwiggyGoLangProject/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func   Delete(ctx *gin.Context)int{
	id_s,_:=ctx.GetQuery("expertid")
	id,_ := primitive.ObjectIDFromHex(id_s)
	return service.Delete(id)
}

func  SignUpExpert(username string,skill string, email string, location int)entity.Expert{
	return service.SignUpExpert(username ,skill , email, location )
}

func GetAllExperts(ctx *gin.Context)[]entity.Expert{
	return service.GetAllExperts()
}

func   WorkDone(ctx *gin.Context){
	a,_:=ctx.GetQuery("userid")
	b,_:=ctx.GetQuery("expertid")
    a_s,_:=primitive.ObjectIDFromHex(a)
	b_s,_:=primitive.ObjectIDFromHex(b)
	fmt.Println(a_s,b_s)
	service.WorkDone(a_s,b_s);
}

func  BookEmployee(ctx *gin.Context) (entity.Expert,int){
	skill,_:=ctx.GetQuery("skill")
	id_s,_:=ctx.GetQuery("userid")
	id,_:=primitive.ObjectIDFromHex(id_s)
	return service.BookEmployee(skill,id)
}

func  AddRating(ctx *gin.Context){
	var review entity.RatingStruct
	ctx.BindJSON(&review)
	id_s,_:=ctx.GetQuery("expertid")
	// id,_ := strconv.ParseInt(id_s,10,8)
	id,_:=primitive.ObjectIDFromHex(id_s)
 	service.AddRating(review.Rating,review.Review,id)
}

func   GetExperts(ctx *gin.Context)[]entity.Expert{
	skill,_ := ctx.GetQuery("skill")
	return service.GetExperts(skill)
}

type Temp struct{
	Username string `json:"name"`
	Skill string `json:"skill"`
	Email string `json:"email"`
	Location int `json:"location"`
}

func  DirectSignUp(ctx *gin.Context)entity.Expert{
	var expert Temp
	ctx.BindJSON(&expert)
	fmt.Println(expert)
	return service.SignUpExpert(expert.Username,expert.Skill,expert.Email,expert.Location)
}

func   GetExpertByID(ctx *gin.Context)entity.Expert{
	id_s,_:=ctx.GetQuery("expertid")
	id,_:=primitive.ObjectIDFromHex(id_s)
	return service.GetExpertByID(id)
}

func  FilterExpert(ctx *gin.Context) []entity.Expert{
	skill,_:=ctx.GetQuery("skill")
	id_s,_:=ctx.GetQuery("rating")
	rating,_ := strconv.ParseInt(id_s,10,8)
	return service.FilterExpert(skill,int(rating))
}

func GetSkills()[]string{
	return service.GetSkills()
}