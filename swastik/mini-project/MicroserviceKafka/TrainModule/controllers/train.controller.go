package controllers

import (
	"net/http"
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/swastiksahoo153/train-module/models"
	"github.com/swastiksahoo153/train-module/services"
)

type TrainController struct{
	TrainService services.TrainService
}

func New(trainService services.TrainService) TrainController{
	return TrainController{
		TrainService: trainService,
	}
}

func (tc *TrainController) CreateTrain(ctx *gin.Context) {
	var train models.Train
	if err := ctx.ShouldBindJSON(&train); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return 
	}
	err := tc.TrainService.CreateTrain(&train)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK,  gin.H{"message": "success"})
}

func (tc *TrainController) GetTrain(ctx *gin.Context) {
	trainnumber := ctx.Param("train_number")
	train, err := tc.TrainService.GetTrain(&trainnumber)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK,  train)
}

func (tc *TrainController) GetAll(ctx *gin.Context) {
	trains, err := tc.TrainService.GetAll()
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, trains)
}

func (tc *TrainController) UpdateTrain(ctx *gin.Context) {
	var train models.Train
	if err := ctx.ShouldBindJSON(&train); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return 
	}
	err := tc.TrainService.UpdateTrain(&train)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK,  gin.H{"message": "success"})
}

func (tc *TrainController) DeleteTrain(ctx *gin.Context) {
	trainnumber := ctx.Param("train_number")
	err := tc.TrainService.DeleteTrain(&trainnumber)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK,  gin.H{"message": "success"})
}

func (tc *TrainController) RegisterTrainRoutes(rg *gin.RouterGroup) {
	trainroute := rg.Group("/train")
	trainroute.POST("/create", tc.CreateTrain)
	trainroute.GET("/get/:train_number", tc.GetTrain)
	trainroute.GET("/getall", tc.GetAll)
	trainroute.PATCH ("/update", tc.UpdateTrain)
	trainroute.DELETE("/delete/:train_number", tc.DeleteTrain)
}
