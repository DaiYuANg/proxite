package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"go.uber.org/fx"
	"proxite/module/config"
	"proxite/module/constant"
)

type NewFiberParam struct {
	fx.In
	Cfg    *config.Config
	Engine *html.Engine
}

func newFiber(param NewFiberParam) *fiber.App {
	engine, cfg := param.Engine, param.Cfg
	app := fiber.New(
		fiber.Config{
			Views:                 engine,
			PassLocalsToViews:     true,
			DisableStartupMessage: false,
			StrictRouting:         true,
			ServerHeader:          constant.AppName,
			AppName:               constant.AppName,
			WriteBufferSize:       cfg.WriteBufferSize,
			Prefork:               cfg.Prefork,
		},
	)
	return app
}
