package route;
import ( 
	"github.com/gin-gonic/gin"
	"github.com/Udaysonu/SwiggyGoLangProject/controller"
	"github.com/Udaysonu/SwiggyGoLangProject/middleware"
 )
 

func UserRouter(g *gin.RouterGroup){
	 
	g.POST("/signuser",func(ctx *gin.Context){
		
		ctx.JSON(200,gin.H{"message":controller.SignUpUser(ctx)})
	})

	g.POST("/isuserpresent",func(ctx *gin.Context){
		ctx.JSON(200,controller.IsUserPresent(ctx))
	})

	g.POST("/getuser",func(ctx *gin.Context){
		ctx.JSON(200,controller.GetUser(ctx))
	})

	g.POST("/loginuser",func(ctx *gin.Context){
		result:=middleware.Login(ctx)
		ctx.JSON(200,gin.H{"token":result})
	})


}

func GetSkills()[]string{
	return []string{"painter","plumber","carpenter"}
}