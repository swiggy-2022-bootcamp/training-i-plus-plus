package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/practice/entity"
	"github.com/practice/service"
)
type VideoController interface{
	FindAll() []entity.Video
	Save(ctx *gin.Context) (entity.Video,error)
}

type controller struct{
	service service.VideoService
}

func New(service service.VideoService) VideoController{
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll()[]entity.Video{
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) (entity.Video , error){
	var video entity.Video
	err:=ctx.ShouldBindJSON(&video)
	if err!=nil{
		return entity.Video{},err
	}
	fmt.Println(video)
	c.service.Save(video)
	return video,nil
}