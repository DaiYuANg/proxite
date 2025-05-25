package config

import (
	"context"
	"github.com/knadh/koanf/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("config",
	fx.Provide(
		newKoanf,
		fx.Annotate(
			load,
			fx.ParamTags(``, `name:"path"`),
		),
	),
	fx.Invoke(
		printer,
	),
)

func newKoanf() *koanf.Koanf {
	return koanf.New(".")
}

func printer(lc fx.Lifecycle, logger *zap.SugaredLogger, k *koanf.Koanf) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info(k.Sprint())
			return nil
		},
	})
}
