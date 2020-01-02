package route

import (
	"boxDemo/controller"
	logicApi "boxDemo/logic/api"
	"boxDemo/model"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	var boxModel model.Box
	boxLogic := logicApi.NewBoxLogic(&boxModel)
	boxController := controller.NewBoxController(boxLogic)
	boxRouteGroup := router.Group("/box")
	{
		boxRouteGroup.GET("/list", boxController.GetBoxList)
	}
}
