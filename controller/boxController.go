package controller

import (
	"boxDemo/logic/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	BoxController struct {
		BoxLogic *logicApi.BoxLogic
	}
)

func NewBoxController(boxLogic *logicApi.BoxLogic) *BoxController {
	return &BoxController{BoxLogic: boxLogic}
}

func (c *BoxController) GetBoxList(ctx *gin.Context) {
	r := new(logicApi.BoxListRequest)
	if err := ctx.ShouldBindJSON(r); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	res, err := c.BoxLogic.GetBoxList(r)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": res})
}
