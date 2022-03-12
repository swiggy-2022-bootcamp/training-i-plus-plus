package controller

import (
	"fmt"
	"github.com/Udaysonu/SwiggyGoLangProject/entity"
	"github.com/Udaysonu/SwiggyGoLangProject/service"
	"github.com/gin-gonic/gin"
	"strconv"
)
 type ExpertController interface{
	GetSkills() []string
	WorkDone(ctx *gin.Context)
	// AddRating()
	SignUpExpert(username string,skill string, email string)
	BookEmployee(ctx *gin.Context) (entity.Expert,int)
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

	fmt.Println(skill)
	return c.service.BookEmployee(skill,int(id))
}