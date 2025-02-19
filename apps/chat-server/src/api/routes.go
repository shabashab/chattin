package api

import (
	"github.com/shabashab/chattin/apps/chat-server/src/api/controllers"
	"github.com/shabashab/chattin/apps/chat-server/src/api/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type setupRoutesProvide struct {
	fx.In

	Router *gin.Engine

	HealthController *controllers.HealthController
	AuthController   *controllers.AuthController

	AuthMiddleware *middleware.AuthMiddleware
}

func setupRoutes(p setupRoutesProvide) {
	api := p.Router.Group("/api/v1/")

	api.GET("/health", p.HealthController.GetHealth)
	api.POST("/auth/debug/login", p.AuthController.DebugLogin)

	authorized := api.Group("/", p.AuthMiddleware.Handler)
	{
		auth := authorized.Group("/auth")
		{
			auth.GET("/iam", p.AuthController.GetCurrentUser)
		}
	}

}
