package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shabashab/chattin/apps/chat-server/src/api/dtos"
	"github.com/shabashab/chattin/apps/chat-server/src/services"
)

type AuthController struct {
	jwtService *services.JwtService
}

type authenticateDebugBody struct {
	UserId uint `json:"userId"`
}

func NewAuthController(jwtService *services.JwtService) (*AuthController) {
	return &AuthController{
		jwtService: jwtService,
	}
}

func (c AuthController) GetCurrentUser(ctx *gin.Context) {
	authenticatedUser := GetRequestAuthenticatedUserOrPanic(ctx)

	ctx.JSON(http.StatusOK, dtos.NewIamDto(authenticatedUser))
}

func (c AuthController) DebugLogin(ctx *gin.Context) {
	body := &authenticateDebugBody{}

	err := ctx.BindJSON(&body)

	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := c.jwtService.CreateTokenForUserId(body.UserId)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dtos.NewDebugLoginDto(token))
}