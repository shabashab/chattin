package controllers

import (
	"chattin/chat-server/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
	healthService *services.HealthService
}

func NewHealthController(healthService *services.HealthService) (*HealthController) {
	return &HealthController{
		healthService: healthService,
	}
}

func (controller HealthController) GetHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": controller.healthService.GetHealthStatus(),
	})
}