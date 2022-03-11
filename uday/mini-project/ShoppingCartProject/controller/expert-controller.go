package controller

import (
	"fmt"

	"github.com/Udaysonu/SwiggyGoLangProject/entity"
	"github.com/Udaysonu/SwiggyGoLangProject/service"
	"github.com/gin-gonic/gin"
)

type ExpertController interface{
	GetSkills() []string
	// WorkDone(int,int)
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

func (c *controller) BookEmployee(ctx *gin.Context) (entity.Expert,int){
	skill,_:=ctx.GetQuery("skill")
	fmt.Println(skill)
	return c.service.BookEmployee(skill,1)
}