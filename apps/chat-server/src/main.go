package main

import (
	"chattin/chat-server/src/api"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		api.Module,

		fx.Provide(
			zap.NewProduction,
		),
	).Run()
}