package routes

import(
    // controller "github.com/bhatiachahat/mongoapi/controllers"
	"bhatiachahat/user-service/middleware"
	controller "bhatiachahat/user-service/controller"
	 "github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine){
	//incomingRoutes.Use(middleware.Authenticate())
	// incomingRoutes.GET("/users", controller.GetUsers())
	// incomingRoutes.GET("/users/:user_id", controller.GetUser())

	public := router.Group("")
	public.POST("/users/signup", controller.Signup())
	public.POST("/users/login", controller.Login())
	// public.POST("/generalUserRegistration",controllers.RegisterGeneralUser())
	// public.POST("/generalUserLogin",controllers.LoginGeneralUser())
	  private := router.Group("")
	 private.Use(middleware.Authenticate())

	 private.POST("/users/addtocart/:user_id", controller.AddToCart())
	 private.GET("/users/getcart/:user_id", controller.GetCart())


	// private.GET("/generalUser/:id",controllers.GetGeneralUserByID())
	// private.PUT("/generalUser/:id", controllers.EditGeneralUserByID())
	// private.DELETE("/generalUser/:id", controllers.DeleteGeneralUserByID())
	// private.GET("/generalUsers", controllers.GetAllGeneralUsers())
	

}