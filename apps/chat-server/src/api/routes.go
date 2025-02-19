package api

import (
	"github.com/shabashab/chattin/apps/chat-server/src/api/controllers"
	"github.com/shabashab/chattin/apps/chat-server/src/api/middleware"

	_ "github.com/shabashab/chattin/apps/chat-server/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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

	p.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
