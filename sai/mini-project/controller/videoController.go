package controller

import (
	"cherie.com/golang-gin/customValidators"
	"cherie.com/golang-gin/entity"
	"cherie.com/golang-gin/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	Save(ctx *gin.Context) error
	FindAll() []entity.Video
}

type videoController struct {
	service service.VideoService
}

var validateIsCool *validator.Validate

func New(service service.VideoService) VideoController {
	validateIsCool = validator.New()
	validateIsCool.RegisterValidation("is-cool", customValidators.ValidateCoolTitle)
	return &videoController{
		service: service,
	}
}

func (controller *videoController) Save(ctx *gin.Context) error {

	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validateIsCool.Struct(video)
	if err != nil {
		return err
	}
	controller.service.Save(video)
	return nil
}

func (controller *videoController) FindAll() []entity.Video {
	return controller.service.FindAll()
}
