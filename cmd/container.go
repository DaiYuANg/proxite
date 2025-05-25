package cmd

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"proxite/module/config"
	"proxite/module/logger"
	"proxite/module/server"
)

func container(configfile string) *fx.App {
	return fx.New(
		fx.Provide(
			fx.Annotate(
				func() string {
					return configfile
				},
				fx.ResultTags(`name:"path"`),
			),
		),
		logger.Module,
		config.Module,
		server.Module,
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)
}
