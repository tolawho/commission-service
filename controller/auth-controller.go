package controller

import (
	"github.com/gin-gonic/gin"
	"medici.vn/commission-serivce/dto"
	"medici.vn/commission-serivce/helper"
	"medici.vn/commission-serivce/models"
	"medici.vn/commission-serivce/services"
	"net/http"
)

// AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtService  services.JWTService
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(authService services.AuthService, jwtService services.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if user, ok := authResult.(models.User); ok {
		generatedToken := c.jwtService.GenerateToken(uint(user.ID))
		user.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", user)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}
