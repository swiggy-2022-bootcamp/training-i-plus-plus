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
		boolean:=middleware.CheckAuth(ctx)
		if boolean==false{
			ctx.JSON(404,gin.H{"message":"Token error, Maybe token expired!"})
		} else{
		ctx.JSON(200,controller.GetSkills())
		}
	})

	g.DELETE("/delete",func(ctx *gin.Context){
		result:=controller.Delete(ctx)
		if result==200{
			ctx.JSON(200,gin.H{"Deleted":"Successfully Deleted"})
		} else{
			ctx.JSON(500,gin.H{"Deleted":"Error in deleting"})
		}
	})

	g.GET("/get",func(ctx *gin.Context){
		boolean:=middleware.CheckAuth(ctx)
		if boolean==false{
			ctx.JSON(404,gin.H{"message":"Token error! Please check token"})
		} else {
			result,err:=controller.BookEmployee(ctx)
			if err==200{
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
			result:=controller.GetAllExperts(ctx)
			ctx.JSON(200,result)	
		}
	})

	g.GET("/done",func(ctx *gin.Context){
		boolean:=middleware.CheckAuth(ctx)
		if boolean==false{
			ctx.JSON(404,gin.H{"message":"Token error! Please check token"})
		} else {
			controller.WorkDone(ctx)
			ctx.JSON(200,gin.H{"message":"Work completed"})		
		}
	})


	g.GET("/experts",func(ctx *gin.Context){
		boolean:=middleware.CheckAuth(ctx)
		if boolean==false{
			ctx.JSON(404,gin.H{"message":"Token error! Please check token"})
		} else {
			ctx.JSON(200,gin.H{"data":controller.GetExperts(ctx)})	
		}
	})


	g.POST("/addrating",func(ctx *gin.Context){
		boolean:=middleware.CheckAuth(ctx)
		if boolean==false{
			ctx.JSON(404,gin.H{"message":"Token error! Please check token"})
		} else {
		controller.AddRating(ctx)
		ctx.JSON(200,gin.H{"message":"Rating Added"})
		}
	})


	g.GET("/getexpert",func(ctx *gin.Context){
		ctx.JSON(200,controller.GetExpertByID(ctx))
	})


	g.POST("/signexpert",func(ctx *gin.Context){
		ctx.JSON(200,gin.H{"message":controller.DirectSignUp(ctx)})
	})

	g.GET("/filter",func(ctx *gin.Context){
		ctx.JSON(200,controller.FilterExpert(ctx))
	})

}