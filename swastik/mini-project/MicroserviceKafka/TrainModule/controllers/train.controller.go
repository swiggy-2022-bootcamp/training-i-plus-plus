package controllers

import (
	"net/http"
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/swastiksahoo153/MicroserviceKafka/TrainModule/models"
	"github.com/swastiksahoo153/MicroserviceKafka/TrainModule/services"
	"github.com/swastiksahoo153/MicroserviceKafka/TrainModule/middlewares"
)

type TrainController struct{
	TrainService services.TrainService
}

func New(trainService services.TrainService) TrainController{
	return TrainController{
		TrainService: trainService,
	}
}


// @Summary Register Train 
// @Description To register a new train for the app.
// @Tags Train
// @Schemes
// @Accept json
// @Produce json
// @Param        train	body	models.Train  true  "Train structure"
// @Success	201  {string} 	success
// @Failure	400  {number} 	http.http.StatusBadRequest
// @Failure	502  {number} 	http.StatusBadGateway
// @Security Bearer Token
// @Router /train/register [POST]
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


// @Summary Get Train
// @Description To get train details.
// @Tags Train
// @Schemes
// @Accept json
// @Param train_number path string true "Train Number"
// @Produce json
// @Success	200  {object} 	models.Train
// @Failure	502  {number} 	http.StatusBadGateway
// @Security Bearer Token
// @Router /train/get/{train_number} [GET]
func (tc *TrainController) GetTrain(ctx *gin.Context) {
	trainnumber := ctx.Param("train_number")
	train, err := tc.TrainService.GetTrain(&trainnumber)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK,  train)
}


// @Summary Get all Train details
// @Description To get every train detail.
// @Tags Train
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {array} 	models.Train
// @Failure	502  {number} 	http.StatusBadGateway
// @Security Bearer Token
// @Router /train/getall [GET]
func (tc *TrainController) GetAll(ctx *gin.Context) {
	trains, err := tc.TrainService.GetAll()
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, trains)
}


// @Summary Update Train
// @Description To update train details.
// @Tags Train
// @Schemes
// @Accept json
// @Param train_number path string true "Train Number"
// @Produce json
// @Success	200  {string} 	success
// @Failure	502  {number} 	http.StatusBadGateway
// @Security Bearer Token
// @Router /train/update/{train_number} [PATCH]
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


// @Summary Delete Train
// @Description To remove a particular train.
// @Tags Train
// @Schemes
// @Accept json
// @Param train_number path string true "Train Number"
// @Produce json
// @Success	200  {string} 	success
// @Failure	502  {number} 	http.StatusBadGateway
// @Security Bearer Token
// @Router /train/delete/{train_number} [DELETE]
func (tc *TrainController) DeleteTrain(ctx *gin.Context) {
	trainnumber := ctx.Param("train_number")
	err := tc.TrainService.DeleteTrain(&trainnumber)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK,  gin.H{"message": "success"})
}


//Routes
func (tc *TrainController) RegisterTrainRoutes(rg *gin.RouterGroup) {
	trainroute := rg.Group("")
	trainroute.Use(middlewares.AuthenticateJWT())
	trainroute.POST("/train/register", tc.CreateTrain)
	trainroute.GET("/train/get/:train_number", tc.GetTrain)
	trainroute.GET("/train/getall", tc.GetAll)
	trainroute.PATCH ("/train/update", tc.UpdateTrain)
	trainroute.DELETE("/train/delete/:train_number", tc.DeleteTrain)
}
