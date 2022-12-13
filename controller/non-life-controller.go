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
	Temporary(ctx *gin.Context)
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
	//f, err := os.OpenFile("commission.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	////defer to close when you're done with it, not because you think it's idiomatic!
	//defer func(f *os.File) {
	//	err := f.Close()
	//	if err != nil {
	//		return
	//	}
	//}(f)
	//
	////set output of logs to f
	//log.SetOutput(f)
	//
	////test case
	//log.Println("The URL: ", ctx.Request.Host+ctx.Request.URL.Path)

	id, err := strconv.ParseUint(ctx.Param("contract_id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	_, err = n.pntDailyCommissionService.Calculator(uint(id))
	if err != nil {
		response := helper.BuildErrorResponse("Not found!", "", err)
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	response := helper.BuildResponse(true, "OK!", true)
	ctx.JSON(http.StatusOK, response)
	return
}

func (n nonLifeController) Temporary(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("contract_id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	err = n.pntDailyCommissionService.Temporary(uint(id))
	if err != nil {
		response := helper.BuildErrorResponse("Error", "", err)
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	response := helper.BuildResponse(true, "OK!", true)
	ctx.JSON(http.StatusOK, response)
	return
}
