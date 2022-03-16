package controller

import (
	"fmt"
	"strconv"
	"github.com/Udaysonu/SwiggyGoLangProject/entity"
	"github.com/Udaysonu/SwiggyGoLangProject/service"
	"github.com/gin-gonic/gin"
	
)
type ExpertController interface{
	GetSkills() []string
	WorkDone(ctx *gin.Context)
	AddRating(ctx *gin.Context)
	SignUpExpert(username string,skill string, email string)
	BookEmployee(ctx *gin.Context) (entity.Expert,int)
	GetExperts(ctx *gin.Context) []entity.Expert
	DirectSignUp(ctx *gin.Context)
	GetExpertByID(ctx *gin.Context) entity.Expert
	FilterExpert(ctx *gin.Context) []entity.Expert
}

type controller struct{
	service service.ExpertService
}

func NewExpertController(service service.ExpertService) ExpertController{	
	return &controller{
		service,
	}
}


func (c *controller) GetSkills()[]string{
	return c.service.GetSkills()
}

func (c *controller) SignUpExpert(username string,skill string, email string){
	c.service.SignUpExpert(username ,skill , email )
}

func (c *controller) WorkDone(ctx *gin.Context){
	b,_:=ctx.GetQuery("userid")
	a,_:=ctx.GetQuery("expertid")
	a_i,_ := strconv.ParseInt(a,10,4)
    b_i,_:=strconv.ParseInt(b,10,4)

	c.service.WorkDone(int(a_i),int(b_i));
}

func (c *controller) BookEmployee(ctx *gin.Context) (entity.Expert,int){
	skill,_:=ctx.GetQuery("skill")
	id_s,_:=ctx.GetQuery("userid")
	id,_ := strconv.ParseInt(id_s,10,4)
	return c.service.BookEmployee(skill,int(id))
}

func (c *controller)AddRating(ctx *gin.Context){
	var review entity.RatingStruct
	ctx.BindJSON(&review)
	id_s,_:=ctx.GetQuery("expertid")
	id,_ := strconv.ParseInt(id_s,10,8)
 	c.service.AddRating(review.Rating,review.Review,int(id))
}

func (c *controller) GetExperts(ctx *gin.Context)[]entity.Expert{
	skill,_ := ctx.GetQuery("skill")
	return c.service.GetExperts(skill)
}
type Temp struct{
	Username string `json:"name"`
	Skill string `json:"skill"`
	Email string `json:"email"`
}
func (c *controller) DirectSignUp(ctx *gin.Context){
	var expert Temp
	ctx.BindJSON(&expert)
	fmt.Println(expert)
	c.service.SignUpExpert(expert.Username,expert.Skill,expert.Email)
}

func (c *controller) GetExpertByID(ctx *gin.Context)entity.Expert{
	id_s,_:=ctx.GetQuery("expertid")
	id,_ := strconv.ParseInt(id_s,10,8)
	return c.service.GetExpertByID(int(id))
}

func (c *controller) FilterExpert(ctx *gin.Context) []entity.Expert{
	skill,_:=ctx.GetQuery("skill")
	id_s,_:=ctx.GetQuery("rating")
	rating,_ := strconv.ParseInt(id_s,10,8)
	return c.service.FilterExpert(skill,int(rating))
}

func GetSkills()[]string{
	return []string{"painter","plumber","carpenter"}
}