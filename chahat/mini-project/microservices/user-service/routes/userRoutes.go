package routes

import(
  
	"bhatiachahat/user-service/middleware"
	controller "bhatiachahat/user-service/controller"
	 "github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine){
	

	public := router.Group("")
	public.POST("/users/signup", controller.Signup())
	public.POST("/users/login", controller.Login())
	
	  private := router.Group("")
	 private.Use(middleware.Authenticate())

	 private.POST("/users/addtocart/:user_id", controller.AddToCart())
	 private.GET("/users/getcart/:user_id", controller.GetCart())


	 private.GET("/users/:id",controller.GetUser())
	 private.PUT("/users/:id", controller.EditUser())
	 private.DELETE("/users/:id", controller.DeleteUser())
	 private.GET("/users", controller.GetAllUsers())
	

}