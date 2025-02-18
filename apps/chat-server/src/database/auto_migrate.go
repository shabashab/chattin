package database

import (
	"chattin/chat-server/src/database/models"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

type AutoMigrateParams struct {
	fx.In

	DB *gorm.DB
	Models []models.Model `group:"models"`
}

func autoMigrate(p AutoMigrateParams) (error) {
	for _, model := range p.Models {
		err := p.DB.AutoMigrate(model)

		if(err != nil) {
			return err
		}
	}

	return nil
}