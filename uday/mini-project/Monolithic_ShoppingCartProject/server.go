package main

import ( 
	"github.com/gin-gonic/gin"
 	// "github.com/Udaysonu/SwiggyGoLangProject/config"
	// "github.com/Udaysonu/SwiggyGoLangProject/service"
 	"github.com/Udaysonu/SwiggyGoLangProject/route"
	"fmt"
)

 

func main(){

	// expertService.InitDB()

	server:=gin.Default()

	expertRoute:=server.Group("/expert")
	{
		route.ExpertRouter(expertRoute.Group("/"))
	}

	userRoute:=server.Group("/user")
	{
		route.UserRouter(userRoute.Group("/"))
	}

	server.Run(":8080")
	fmt.Println("Server started Listening at port:8080")
}
