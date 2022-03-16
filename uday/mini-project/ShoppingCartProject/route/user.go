package route;
import ( 
	"github.com/gin-gonic/gin"
	"github.com/Udaysonu/SwiggyGoLangProject/controller"
	"github.com/Udaysonu/SwiggyGoLangProject/config"
	"github.com/Udaysonu/SwiggyGoLangProject/service"
	"github.com/Udaysonu/SwiggyGoLangProject/middleware"
 )
var ( 
	expertService service.ExpertService=config.GetexpertService()
	expertController controller.ExpertController=config.GetexpertController()
	userService service.UserService=service.UserNew()
	userController controller.UserController=controller.UserNewController(userService)
)


func UserRouter(g *gin.RouterGroup){
	 
	g.POST("/signuser",func(ctx *gin.Context){
		userController.SignUpUser(ctx)
		ctx.JSON(200,gin.H{"message":"Sign In Successful"})
	})

	g.POST("/isuserpresent",func(ctx *gin.Context){
		ctx.JSON(200,userController.IsUserPresent(ctx))
	})

	g.POST("/getuser",func(ctx *gin.Context){
		ctx.JSON(200,userController.GetUser(ctx))
	})

	g.POST("/loginuser",func(ctx *gin.Context){
		result:=middleware.Login(ctx,userService)
		ctx.JSON(200,gin.H{"token":result})
	})


}

func GetSkills()[]string{
	return []string{"painter","plumber","carpenter"}
}