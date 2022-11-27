package controller

import (
	"github.com/gin-gonic/gin"
	"medici.vn/commission-serivce/helper"
	"medici.vn/commission-serivce/services"
	"net/http"
	"strconv"
)

// NonLifeController interface is a contract what this controller can do
type NonLifeController interface {
	Calculator(ctx *gin.Context)
}

type nonLifeController struct {
	pntDailyCommissionService services.PntDailyCommissionService
}

// NewNonLifeController creates a new instance of NonLifeController
func NewNonLifeController(pntDailyCommissionService services.PntDailyCommissionService) NonLifeController {
	return &nonLifeController{
		pntDailyCommissionService: pntDailyCommissionService,
	}
}

func (n nonLifeController) Calculator(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("contract_id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	commission := n.pntDailyCommissionService.Calculator(uint(id))
	response := helper.BuildResponse(true, "OK!", commission)
	ctx.JSON(http.StatusOK, response)
	return

}
