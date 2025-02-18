package database

import (
	"chattin/chat-server/src/config/configs"
	"chattin/chat-server/src/database/models"

	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Module = fx.Module("database",
	models.Module,

	fx.Provide(
		newDatabase,
	),

	fx.Invoke(
		autoMigrate,
	),
)

func newDatabase(dbConfig *configs.DBConfig, ) (*gorm.DB, error) {
	pg := postgres.Open(dbConfig.ConnectionUrl)

	db, err := gorm.Open(pg)

	if (err != nil) {
		return nil, err
	}

	return db, nil
}
