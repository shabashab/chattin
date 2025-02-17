package api

import (
	"chattin/chat-server/src/api/controllers"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type setupRoutesProvide struct {
	fx.In

	Router *gin.Engine

	HealthController *controllers.HealthController
}

func setupRoutes(p setupRoutesProvide) {
	apiGroup := p.Router.Group("/api/v1/")

	{
		apiGroup.GET("/health", p.HealthController.GetHealth)
	}
}