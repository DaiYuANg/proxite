package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"proxite/module/config"
)

func proxyMiddleware(cfg *config.Config, app *fiber.App, log *zap.SugaredLogger) {
	lo.ForEach(cfg.SpaProxies, func(spa config.SpaProxy, index int) {
		lo.ForEach(spa.Proxy, func(rule config.ProxyRule, index int) {
			pathPrefix := spa.Root + rule.PathPrefix
			target := rule.Target
			proxyPath := pathPrefix + "*"
			targetCopy := target

			app.All(proxyPath, func(ctx *fiber.Ctx) error {
				path := targetCopy + ctx.Path()
				log.Debugf("proxy target:%v", path)
				return proxy.Do(ctx, path)
			})
		})
		app.Static(spa.Root+"*", spa.SpaPath, fiber.Static{
			ByteRange: true,
			Compress:  true,
		})
	})

	//app.Use(redirect.New(redirect.Config{
	//	Rules: map[string]string{
	//		"/old":   "/new",
	//		"/old/*": "/new/$1",
	//	},
	//	StatusCode: 301,
	//}))
	//
	//app.Use(rewrite.New(rewrite.Config{
	//	Rules: map[string]string{
	//		"/old":   "/new",
	//		"/old/*": "/new/$1",
	//	},
	//}))
}
