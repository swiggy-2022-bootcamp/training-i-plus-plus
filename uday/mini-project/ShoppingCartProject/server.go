package main

import ( 
	"github.com/gin-gonic/gin"
	"github.com/Udaysonu/SwiggyGoLangProject/controller"
	"github.com/Udaysonu/SwiggyGoLangProject/config"
	"github.com/Udaysonu/SwiggyGoLangProject/service"
	"github.com/Udaysonu/SwiggyGoLangProject/middleware"

)

var (
	UES service.UserExpertService=config.GetUES()
	UEC controller.UserExpertController=config.GetUEC()
	expertService service.ExpertService=config.GetexpertService()
	expertController controller.ExpertController=config.GetexpertController()
	userService service.UserService=service.UserNew()
	userController controller.UserController=controller.UserNewController(userService)
)


func main(){

	expertService.InitDB()

	server:=gin.Default()

	server.GET("/services", func(ctx *gin.Context){
		boolean:=middleware.CheckAuth(ctx,userService)
		if boolean==false{
			ctx.JSON(404,gin.H{"message":"Token error, Maybe token expired!"})
		} else{
		ctx.JSON(200,expertController.GetSkills())
		}
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

	server.GET("/experts",func(ctx *gin.Context){
		ctx.JSON(200,expertController.GetExperts(ctx))
	})

	server.POST("/addrating",func(ctx *gin.Context){
		expertController.AddRating(ctx)
		ctx.JSON(200,gin.H{"message":"Rating Added"})
	})

	server.GET("/getexpert",func(ctx *gin.Context){
		ctx.JSON(200,expertController.GetExpertByID(ctx))
	})

	server.POST("/signexpert",func(ctx *gin.Context){
		expertController.DirectSignUp(ctx)
		ctx.JSON(200,gin.H{"message":"Sign In Successful"})
	})

	server.GET("/filter",func(ctx *gin.Context){
		ctx.JSON(200,expertController.FilterExpert(ctx))
	})
 
	server.POST("/signuser",func(ctx *gin.Context){
		userController.SignUpUser(ctx)
		ctx.JSON(200,gin.H{"message":"Sign In Successful"})
	})

	server.POST("/isuserpresent",func(ctx *gin.Context){
		ctx.JSON(200,userController.IsUserPresent(ctx))
	})

	server.POST("/getuser",func(ctx *gin.Context){
		ctx.JSON(200,userController.GetUser(ctx))
	})

	server.POST("/loginuser",func(ctx *gin.Context){
		result:=middleware.Login(ctx,userService)
		ctx.JSON(200,gin.H{"token":result})
	})

	server.Run(":8080")
}