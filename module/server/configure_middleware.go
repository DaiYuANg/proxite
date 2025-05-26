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
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"proxite/module/config"
	"proxite/module/constant"
)

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_requests_total",
			Help: "Total number of requests",
		},
		[]string{"path"},
	)
)

func configureLoggerMiddleware(cfg *config.Config, app *fiber.App) {
	app.Use(logger.New())
}

func configureOtelFiberConfig(cfg *config.Config, app *fiber.App) {
	app.Use(otelfiber.Middleware())
}

func configureZapLogger(cfg *config.Config, app *fiber.App, logger *zap.Logger) {
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))
}

func configureCors(cfg *config.Config, app *fiber.App) {
	app.Use(cors.New())
}

func configureMonitor(cfg *config.Config, app *fiber.App) {
	app.Get("/metrics", monitor.New())
}

func configureHealthcheck(cfg *config.Config, app *fiber.App) {
	app.Use(healthcheck.New())
}

func configureHelmet(cfg *config.Config, app *fiber.App) {
	app.Use(helmet.New())
}

func configureCompress(cfg *config.Config, app *fiber.App) {
	app.Use(compress.New())
}

func configureFavicon(cfg *config.Config, app *fiber.App) {
	app.Use(favicon.New())
}

func configurePprof(cfg *config.Config, app *fiber.App) {
	app.Use(pprof.New(pprof.Config{Prefix: "/debug"}))
}

func configureRequestId(cfg *config.Config, app *fiber.App) {
	app.Use(requestid.New())
}

func configurePrometheus(cfg *config.Config, app *fiber.App) {
	pro := fiberprometheus.New(constant.AppName)
	prometheus.MustRegister(requestCounter)
	pro.RegisterAt(app, "/metrics")
	app.Use(pro.Middleware)
}

func configureRecover(cfg *config.Config, app *fiber.App) {
	app.Use(recover.New())
}

func configureDefaultPage(cfg *config.Config, app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		if len(cfg.SpaProxies) == 0 {
			return c.Status(fiber.StatusNotFound).Render("index", fiber.Map{})
		}
		return c.Status(fiber.StatusNotFound).Render(PageNotFound, fiber.Map{
			"title": "Page Not Found",
			"url":   c.OriginalURL(),
		})
	})
}
