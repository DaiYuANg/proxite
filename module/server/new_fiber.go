package server

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/redirect"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/rewrite"
	"github.com/gofiber/template/html/v2"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"proxite/module/config"
	"proxite/module/constant"
)

type NewFiberParam struct {
	fx.In
	Cfg    *config.Config
	Engine *html.Engine `name:"serverTemplateEngine"`
	Logger *zap.Logger
}

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_requests_total",
			Help: "Total number of requests",
		},
		[]string{"path"},
	)
)

func newFiber(param NewFiberParam) *fiber.App {
	engine, cfg, logger := param.Engine, param.Cfg, param.Logger
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
	app.Use(otelfiber.Middleware())
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))
	configureLoggerMiddleware(cfg, app)
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(favicon.New())
	app.Use(healthcheck.New())
	app.Use(helmet.New())
	app.Get("/metrics", monitor.New())
	app.Use(pprof.New(pprof.Config{Prefix: "/debug"}))
	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/old":   "/new",
			"/old/*": "/new/$1",
		},
		StatusCode: 301,
	}))
	app.Use(requestid.New())
	app.Use(rewrite.New(rewrite.Config{
		Rules: map[string]string{
			"/old":   "/new",
			"/old/*": "/new/$1",
		},
	}))

	pro := fiberprometheus.New(constant.AppName)
	prometheus.MustRegister(requestCounter)
	pro.RegisterAt(app, "/metrics")
	app.Use(pro.Middleware)
	app.Use(recover.New())
	for _, spa := range cfg.SpaProxies {
		for _, rule := range spa.Proxy {
			pathPrefix := spa.Root + rule.PathPrefix
			target := rule.Target

			targetCopy := target

			app.Use(pathPrefix, func(c *fiber.Ctx) error {
				return proxy.Do(c, targetCopy)
			})
		}

		app.Static(spa.Root, spa.SpaPath, fiber.Static{
			Browse:   false,
			Compress: true,
		})
	}

	app.Use(func(c *fiber.Ctx) error {
		if len(cfg.SpaProxies) == 0 {
			return c.Status(fiber.StatusNotFound).Render("index", fiber.Map{})
		}
		return c.Status(fiber.StatusNotFound).Render(PageNotFound, fiber.Map{
			"title": "Page Not Found",
			"url":   c.OriginalURL(),
		})
	})
	return app
}
