package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"proxite/module/config"
)

func configureLoggerMiddleware(cfg *config.Config, app *fiber.App) {
	app.Use(logger.New())
}
