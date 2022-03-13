package main

import (
	// "io"
	// "net/http"
	// "os"
	"github.com/gin-gonic/gin"
	"github.com/Udaysonu/SwiggyGoLangProject/controller"
	"github.com/Udaysonu/SwiggyGoLangProject/config"
	"github.com/Udaysonu/SwiggyGoLangProject/service"

)

var (
	UES service.UserExpertService=config.GetUES()
	UEC controller.UserExpertController=config.GetUEC()
	expertService service.ExpertService=config.GetexpertService()
	expertController controller.ExpertController=config.GetexpertController()
)

// func setupLogOutput(){
// 	f,_:=os.Create("gin.log")
// 	gin.DefaultWriter = io.MultiWriter(f,os.Stdout)
// }

func main(){

	expertService.InitDB()
	server:=gin.Default()
	server.GET("/services", func(ctx *gin.Context){
		ctx.JSON(200,expertController.GetSkills())
	})


	server.GET("/get",func(ctx *gin.Context){
		result,err:=expertController.BookEmployee(ctx)
		if err==200{
			ctx.JSON(200,result)
		} else{
			ctx.JSON(400,gin.H{"message":"All Expers are busy, please try again"})
		}
		
	})

	server.GET("/done",func(ctx *gin.Context){
		expertController.WorkDone(ctx)
		ctx.JSON(200,gin.H{"message":"Work completed"})
	})


	

	// server.POST("/posts",func(ctx *gin.Context){
	// 	result,err:=videoController.Save(ctx)
	// 	if err==nil{
	// 		ctx.JSON(http.StatusBadRequest,result)
	// 	} else{
	// 		ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	// 	}
	// })
	server.Run(":8080")
}