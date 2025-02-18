package database

import (
	"chattin/chat-server/src/config/configs"

	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Module = fx.Module("database",
	fx.Provide(
		NewDatabase,
	),
)

func NewDatabase(dbConfig *configs.DBConfig) (*gorm.DB, error) {
	pg := postgres.Open(dbConfig.ConnectionUrl)

	db, err := gorm.Open(pg)

	if (err != nil) {
		return nil, err
	}

	return db, nil
}