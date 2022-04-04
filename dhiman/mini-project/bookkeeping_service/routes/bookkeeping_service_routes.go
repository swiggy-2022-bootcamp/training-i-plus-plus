package routes

import (
	"github.com/dhi13man/healthcare-app/bookkeeping_service/controllers"
	"github.com/gin-gonic/gin"
)

const BaseURL string = "/bookkeeping";

func GenerateBookKeepingServiceRoutes(router *gin.Engine) {
	bookKeepingRouter := router.Group(BaseURL)
	// Create
	bookKeepingRouter.POST("medicines", controllers.CreateMedicine)
	// Read
	bookKeepingRouter.GET("medicines", controllers.FindAllMedicines)
	bookKeepingRouter.GET("medicines/:id", controllers.FindMedicineByName)
	bookKeepingRouter.GET("medicines/disease/:diseaseName", controllers.FindMedicinesByDiseaseName)
	// Update
	bookKeepingRouter.PUT("medicines/:id", controllers.UpdateMedicines)
	// Delete
	bookKeepingRouter.DELETE("medicines/:id", controllers.DeleteMedicines)
}