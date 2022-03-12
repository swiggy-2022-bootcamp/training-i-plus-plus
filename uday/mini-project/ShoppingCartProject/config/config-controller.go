package config
import (
	"github.com/Udaysonu/SwiggyGoLangProject/service"
	"github.com/Udaysonu/SwiggyGoLangProject/controller"
)
var(
	UES service.UserExpertService=service.NewUserExpertService()
	UEC controller.UserExpertController=controller.NewUserExpert(UES)
	expertService service.ExpertService=service.ExpertNew()
	expertController controller.ExpertController=controller.NewExpertController(expertService)
)

func GetUES() service.UserExpertService{
	return UES
}

func GetUEC() controller.UserExpertController{
	return UEC
}

func GetexpertService() service.ExpertService{
	return expertService
}

func GetexpertController() controller.ExpertController{
	return expertController
}