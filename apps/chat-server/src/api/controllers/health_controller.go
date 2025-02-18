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
	healthStatus := controller.healthService.GetHealthStatus()
	httpStatus := http.StatusOK

	if(!healthStatus.Alive) {
		httpStatus = http.StatusInternalServerError
	}

	ctx.JSON(httpStatus, gin.H{
		"status": healthStatus,
	})
}