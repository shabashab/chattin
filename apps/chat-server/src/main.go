package main

import (
	"chattin/chat-server/src/api"
	"fmt"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			zap.NewProduction(),
		),

		api.Module,

		fx.Invoke(func () {
			fmt.Println("Hello, world!")
		}),
	).Run()
}