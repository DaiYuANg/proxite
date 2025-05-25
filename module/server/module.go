package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"proxite/module/config"
	"proxite/view"
)

var Module = fx.Module("server",
	fx.Provide(
		fx.Annotate(
			templateEngine,
			fx.ResultTags(`name:"serverTemplateEngine"`),
		),
		fx.Annotate(
			newFiber,
			fx.ResultTags(`name:"server"`),
		),
	),
	fx.Invoke(
		startServer,
	),
)

func templateEngine() *html.Engine {
	return html.NewFileSystem(http.FS(view.View), ".html")
}

type StartServerParam struct {
	fx.In
	Lifecycle fx.Lifecycle
	Server    *fiber.App `name:"server"`
	Cfg       *config.Config
	Logger    *zap.SugaredLogger
}

func startServer(param StartServerParam) {
	lc, server, _, log := param.Lifecycle, param.Server, param.Cfg, param.Logger
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					//err := server.Listen(":" + fmt.Sprint(cfg.Port))
					log.Infof("Server started %v", "http://localhost:9876")
					err := server.Listen(":9876")
					if err != nil {
						panic(err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return server.ShutdownWithContext(ctx)
			},
		},
	)
}
