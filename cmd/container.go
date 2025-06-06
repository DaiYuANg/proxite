package cmd

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"proxite/module/cache"
	"proxite/module/config"
	"proxite/module/logger"
	"proxite/module/server"
	"proxite/module/validator"
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
		cache.Module,
		validator.Module,
		logger.Module,
		config.Module,
		server.Module,
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)
}
