package services

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HealthService struct {
	logger *zap.Logger
	db *gorm.DB
}

type HealthStatus struct {
	Alive	bool `json:"alive"`

	Database bool	`json:"database"`
}

func NewHealthService(logger *zap.Logger, db *gorm.DB) (*HealthService) {
	return &HealthService{
		logger: logger,
		db: db,
	}
}

func (healthService HealthService) GetHealthStatus() (HealthStatus) {
	healthService.logger.Info("Health status requested")

	dbHealthStatus := healthService.getDatabaseHealthStatus()

	return HealthStatus{
		Database: dbHealthStatus,
	}
}

func (healthService HealthService) getDatabaseHealthStatus() (bool) {
	db, err := healthService.db.DB()

	if err != nil {
		healthService.logger.Error("database healthcheck failed", zap.Error(err))
		fmt.Println(err)
		return false
	}

	err = db.Ping()

	if (err != nil) {
		healthService.logger.Error("database healthcheck failed", zap.Error(err))
		return false
	}

	return true
}