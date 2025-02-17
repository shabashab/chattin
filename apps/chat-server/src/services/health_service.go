package services

import "go.uber.org/zap"

type HealthService struct {
	logger *zap.Logger
}

func NewHealthService(logger *zap.Logger) (*HealthService) {
	return &HealthService{
		logger: logger,
	}
}

func (healthService HealthService) GetHealthStatus() (string) {
	healthService.logger.Info("Health status requested")

	return "healthy"
}