package api

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/shabashab/chattin/apps/chat-server/src/api/controllers"
	"github.com/shabashab/chattin/apps/chat-server/src/api/middleware"
	"github.com/shabashab/chattin/apps/chat-server/src/config/configs"

	_ "github.com/shabashab/chattin/apps/chat-server/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Module("api",
	controllers.Module,
	middleware.Module,

	fx.Provide(
		newServer,
	),
	fx.Invoke(setupRoutes),
)

// @title           Chatting API
// @version         1.0
// @description     This is the API for the Chatting

// @host      localhost:4000
// @BasePath  /api/v1
func newServer(lc fx.Lifecycle, apiConfig *configs.ApiConfig) (*http.Server, *gin.Engine) {
	router := gin.Default()

	server := &http.Server{
		Addr: apiConfig.Host,
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
