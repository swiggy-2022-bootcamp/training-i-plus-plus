package route
import ( 
	"github.com/gin-gonic/gin"
	 
	"github.com/Udaysonu/SwiggyGoLangProject/middleware"
 
)
func called(ctx *gin.Context)bool{
	return true
}

 

func ExpertRouter(g *gin.RouterGroup){
	
	
	g.GET("/services", func(ctx *gin.Context){
		boolean:=middleware.CheckAuth(ctx,userService)
		if boolean==false{
			ctx.JSON(404,gin.H{"message":"Token error, Maybe token expired!"})
		} else{
		ctx.JSON(200,expertController.GetSkills())
		}
	})


	g.GET("/get",func(ctx *gin.Context){
		result,err:=expertController.BookEmployee(ctx)
		if err==200{
			ctx.JSON(200,result)
		} else{
			ctx.JSON(400,gin.H{"message":"All Expers are busy, please try again"})
		}
	})


	g.GET("/done",func(ctx *gin.Context){
		expertController.WorkDone(ctx)
		ctx.JSON(200,gin.H{"message":"Work completed"})
	})


	g.GET("/experts",func(ctx *gin.Context){
		ctx.JSON(200,expertController.GetExperts(ctx))
	})


	g.POST("/addrating",func(ctx *gin.Context){
		expertController.AddRating(ctx)
		ctx.JSON(200,gin.H{"message":"Rating Added"})
	})


	g.GET("/getexpert",func(ctx *gin.Context){
		ctx.JSON(200,expertController.GetExpertByID(ctx))
	})


	g.POST("/signexpert",func(ctx *gin.Context){
		expertController.DirectSignUp(ctx)
		ctx.JSON(200,gin.H{"message":"Sign In Successful"})
	})

	g.GET("/filter",func(ctx *gin.Context){
		ctx.JSON(200,expertController.FilterExpert(ctx))
	})
}