package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/gofiber/fiber/v2/middleware/redirect"
	"github.com/gofiber/fiber/v2/middleware/rewrite"
	"github.com/samber/lo"
	"proxite/module/config"
)

func proxyMiddleware(cfg *config.Config, app *fiber.App) {
	lo.ForEach(cfg.SpaProxies, func(spa config.SpaProxy, index int) {
		lo.ForEach(spa.Proxy, func(rule config.ProxyRule, index int) {
			pathPrefix := spa.Root + rule.PathPrefix
			target := rule.Target

			targetCopy := target

			app.Use(pathPrefix, func(c *fiber.Ctx) error {
				return proxy.Do(c, targetCopy)
			})
		})
		app.Static(spa.Root, spa.SpaPath, fiber.Static{
			Browse:   false,
			Compress: true,
		})
	})

	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/old":   "/new",
			"/old/*": "/new/$1",
		},
		StatusCode: 301,
	}))

	app.Use(rewrite.New(rewrite.Config{
		Rules: map[string]string{
			"/old":   "/new",
			"/old/*": "/new/$1",
		},
	}))

}
