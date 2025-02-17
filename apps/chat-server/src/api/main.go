package api

import (
	"chattin/chat-server/src/api/controllers"
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func newServer(lc fx.Lifecycle) (*http.Server, *gin.Engine) {
	router := gin.Default()

	server := &http.Server{
		Addr: ":4000",
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			server.Handler = router.Handler()

			listener, err := net.Listen("tcp", server.Addr)

			if err != nil {
				return err
			}

			fmt.Println("Starting http server at", server.Addr)

			go server.Serve(listener)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})

	return server, router
}

var Module = fx.Module("api", 
	controllers.Module,
	fx.Provide(
		newServer,
	),
	fx.Invoke(setupRoutes),
)