package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sing3demons/golanh-api/dto"
	"github.com/sing3demons/golanh-api/entity"
	"github.com/sing3demons/golanh-api/helper"
	"github.com/sing3demons/golanh-api/service"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var form dto.LoginDTO
	if err := ctx.ShouldBind(&form); err != nil {
		resp := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}
	authResult := c.authService.VerifyCredential(form.Email, form.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		resp := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, resp)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (a *authController) Register(ctx *gin.Context) {
	var form dto.RegisterDTO
	if err := ctx.ShouldBind(&form); err != nil {
		resp := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	if !a.authService.IsDuplicateEmail(form.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := a.authService.CreateUser(form)
		token := a.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated,response)
	}
}
