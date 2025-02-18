package main

import (
	"chattin/chat-server/src/api"
	"chattin/chat-server/src/config"
	"chattin/chat-server/src/database"
	"chattin/chat-server/src/services"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		api.Module,
		services.Module,
		config.Module,
		database.Module,

		fx.Provide(
			zap.NewProduction,
		),
	).Run()
}