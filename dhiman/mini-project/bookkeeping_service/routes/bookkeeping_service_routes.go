package routes

import (
	"github.com/dhi13man/healthcare-app/bookkeeping_service/controllers"
	"github.com/gin-gonic/gin"
)

const baseURL string = "/bookkeeping/";

func GenerateBookKeepingServiceRoutes(router *gin.Engine) {
	router.GET(baseURL + "medicines/", controllers.FindAllMedicines)
	router.GET(baseURL + "medicines/:id/", controllers.FindMedicineByName)
	router.GET(baseURL + "medicines/disease/:diseaseName/", controllers.FindMedicinesByDiseaseName)
}