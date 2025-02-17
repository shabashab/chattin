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
	api := p.Router.Group("/api/v1/")

	api.GET("/health", p.HealthController.GetHealth)
}