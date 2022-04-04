package routes

import (
	"github.com/dhi13man/healthcare-app/bookkeeping_service/controllers"
	"github.com/gin-gonic/gin"
)

const baseURL string = "/bookkeeping";

func GenerateBookKeepingServiceRoutes(router *gin.Engine) {
	bookKeepingRouter := router.Group(baseURL)
	bookKeepingRouter.GET("medicines", controllers.FindAllMedicines)
	bookKeepingRouter.GET("medicines/:id", controllers.FindMedicineByName)
	bookKeepingRouter.GET("medicines/disease/:diseaseName", controllers.FindMedicinesByDiseaseName)
}