package route
import ( 
	"github.com/gin-gonic/gin"
	 
	"github.com/Udaysonu/SwiggyGoLangProject/middleware"
	"github.com/Udaysonu/SwiggyGoLangProject/controller"
)
func called(ctx *gin.Context)bool{
	return true
}

func ExpertRouter(g *gin.RouterGroup){
	
	g.GET("/services", func(ctx *gin.Context){
		 
		status,value:=controller.GetSkills()
		if status!=200{
			ctx.JSON(status,value)
		}else{
			ctx.JSON(200,value)
		}

	 
	})

	g.DELETE("/delete/:expertid",func(ctx *gin.Context){
		status:=controller.Delete(ctx)
		if status==200{
			ctx.JSON(200,gin.H{"Deleted":"Successfully Deleted"})
		} else{
			ctx.JSON(500,gin.H{"Deleted":"Error in deleting"})
		}
	})

	g.GET("/getbyskill/:skill",func(ctx *gin.Context){
		boolean:=middleware.CheckAuth(ctx)
		if boolean==false{
			ctx.JSON(404,gin.H{"message":"Token error! Please check token"})
		} else {
			status,result:=controller.GetExperts(ctx)
			if status==200{
				ctx.JSON(200,result)
			} else{
				ctx.JSON(400,gin.H{"message":"All Expers are busy, please try again"})
			}
		}
	})

	g.GET("/get/:skill/:userid",func(ctx *gin.Context){
		boolean:=middleware.CheckAuth(ctx)
		if boolean==false{
			ctx.JSON(404,gin.H{"message":"Token error! Please check token"})
		} else {
			status,result:=controller.BookEmployee(ctx)
			if status==200{
				ctx.JSON(200,result)
			} else{
				ctx.JSON(400,gin.H{"message":"All Expers are busy, please try again"})
			}
		}
	})


	g.GET("/getallexperts",func(ctx *gin.Context){
		boolean:=middleware.CheckAuth(ctx)
		if boolean==false{
			ctx.JSON(404,gin.H{"message":"Token error! Please check token"})
		} else {
			status,result:=controller.GetAllExperts(ctx)
			if status!=200{
				ctx.JSON(status,result)
			}else{
				ctx.JSON(200,result)
			}		}
	})

	g.GET("/done/:userid/:expertid",func(ctx *gin.Context){
		boolean:=middleware.CheckAuth(ctx)
		if boolean==false{
			ctx.JSON(404,gin.H{"message":"Token error! Please check token"})
		} else {
			status,_:=controller.WorkDone(ctx)
			if status!=200{
				ctx.JSON(status,gin.H{"message":"Error encountered"})
			}else{
				ctx.JSON(200,gin.H{"message":"Work Completed"})
			}	
 		}
	})


	g.GET("/experts/:skill",func(ctx *gin.Context){
		boolean:=middleware.CheckAuth(ctx)
		if boolean==false{
			ctx.JSON(404,gin.H{"message":"Token error! Please check token"})
		} else {
			status,result:=controller.GetExperts(ctx)
			if status!=200{
				ctx.JSON(status,result)
			}else{
				ctx.JSON(200,result)
			}		
		}
	})


	g.POST("/addrating/:expertid",func(ctx *gin.Context){
		boolean:=middleware.CheckAuth(ctx)
		if boolean==false{
			ctx.JSON(404,gin.H{"message":"Token error! Please check token"})
		} else {
		controller.AddRating(ctx)
		ctx.JSON(200,gin.H{"message":"Rating Added"})
		}
	})


	g.GET("/getexpert/:expertid",func(ctx *gin.Context){
		status,result:=controller.GetExpertByID(ctx)
		if status!=200{
			ctx.JSON(status,result)
		}else{
			ctx.JSON(200,result)
		}	
	})


	g.POST("/signupexpert",func(ctx *gin.Context){
		status,result:=controller.DirectSignUp(ctx)
		if status!=200{
			ctx.JSON(status,gin.H{"message":result})
		}else{
			ctx.JSON(200,gin.H{"message":result})
		}	
	})

	g.GET("/filter/:skill/:rating",func(ctx *gin.Context){
		status,result:=controller.FilterExpert(ctx)
		if status!=200{
			ctx.JSON(status,result)
		}else{
			ctx.JSON(200,result)
		}	
	})

	g.GET("/waitingreq/:expertid",func(ctx *gin.Context){
		status,result:=controller.GetWaitingRequest(ctx)
		if status!=200{
			ctx.JSON(status,result)
		}else{
			ctx.JSON(200,result)
		}	
	})
	g.GET("/rejectreq/:expertid",func(ctx *gin.Context){
		status,result:=controller.RejectWaitingResult(ctx)
		if status!=200{
			ctx.JSON(status,result)
		}else{
			ctx.JSON(200,result)
		}	
	})
	g.GET("/acceptreq/:expertid",func(ctx *gin.Context){
		status,result:=controller.AcceptWaitingRequest(ctx)
		if status!=200{
			ctx.JSON(status,result)
		}else{
			ctx.JSON(200,result)
		}	
	})
	g.GET("/complete/:cost/:expertid",func(ctx *gin.Context){
		status,result:=controller.CompletedRequest(ctx)
		if status!=200{
			ctx.JSON(status,result)
		}else{
			ctx.JSON(200,result)
		}	
	})
}