package routes

import (
	"github.com/dhi13man/healthcare-app/bookkeeping_service/controllers"
	"github.com/gin-gonic/gin"
)

const BaseURL string = "/bookkeeping";

func GenerateBookKeepingServiceRoutes(router *gin.Engine) {
	bookKeepingRouter := router.Group(BaseURL)
	medicinebookKeepingRouter := bookKeepingRouter.Group("/medicine")
	
	// Create
	medicinebookKeepingRouter.POST("/", controllers.CreateMedicine)
	// Read
	medicinebookKeepingRouter.GET("/", controllers.FindAllMedicines)
	medicinebookKeepingRouter.GET("/:id", controllers.FindMedicineByName)
	medicinebookKeepingRouter.GET("/disease/:diseaseName", controllers.FindMedicinesByDiseaseName)
	// Update
	medicinebookKeepingRouter.PUT("/:id", controllers.UpdateMedicines)
	// Delete
	medicinebookKeepingRouter.DELETE("/:id", controllers.DeleteMedicines)
}