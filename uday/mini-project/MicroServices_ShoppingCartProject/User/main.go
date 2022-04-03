package main

import ( 
	"github.com/gin-gonic/gin"
 	log "github.com/Udaysonu/SwiggyGoLangProject/config"
	// "github.com/Udaysonu/SwiggyGoLangProject/service"
	_ "github.com/Udaysonu/SwiggyGoLangProject/docs"
 	"github.com/Udaysonu/SwiggyGoLangProject/route"
	ginSwagger "github.com/swaggo/gin-swagger"  //gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles"	//swagger embed files
 
)

 
// @title Swagger demo service API
// @version 1.0
// @description This is demo server.
// @termsOfService demo.com

// @contact.name API Support
// @contact.url http://demo.com/support

// @host localhost:8080
// @BasePath /

// @securityDefinitions.basic  BasicAuth

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization


func main(){
	// expertService.InitDB()
	log.InitLogger()
	go log.InitKafka()
	
	server:=gin.Default()

	expertRoute:=server.Group("/expert")
	{
		route.ExpertRouter(expertRoute.Group("/"))
	}

	userRoute:=server.Group("/user")
	{
		route.UserRouter(userRoute.Group("/"))
	}
	server.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run(":8080")
	log.Info.Println("Server started Listening at port:8080")
}
